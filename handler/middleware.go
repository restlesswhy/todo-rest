package handler

import (
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}