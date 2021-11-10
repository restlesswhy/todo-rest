package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todorest "github.com/restlesswhy/todo-rest"
)

type tokenResponce struct {
	AccesToken string	`json:"accesToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) signUp(c *gin.Context) {
	var user todorest.User

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var user UserSignIn

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponce{
		AccesToken: res,
	})
}