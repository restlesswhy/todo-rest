package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/repository"
)

const (
	salt = "alisfn7asfy3y987f34984u34"
	tokenTTL = 12 * time.Hour
)

var signInKey = os.Getenv("SIGNIN_KEY")

type AuthService struct {
	repo repository.Authorization
}

func NewAuthSerice(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todorest.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

type myClaims struct {
	*jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	id, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		StandardClaims: claims,
		UserId: id,
	})


	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) ParseToken(accecTtoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accecTtoken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New(("token clames are not of type *tokenClames"))
	}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}