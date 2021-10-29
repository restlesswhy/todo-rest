package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	splitHeader := strings.Split(header, " ")
	if len(splitHeader) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(splitHeader[1])
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid parse token")
		return
	}

	c.Set("userId", userId)
}

func (h *Handler) GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "userId is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponce(c, http.StatusInternalServerError, "userId have invalid type")
		return 0, errors.New("user have invalid type")
	}

	return idInt, nil
}