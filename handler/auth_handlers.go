package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todorest "github.com/restlesswhy/todo-rest"
)

func (h *Handler) signUp(c *gin.Context) {
	var user todorest.User

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	
}