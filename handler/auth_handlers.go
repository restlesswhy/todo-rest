package handler

import (
	"github.com/gin-gonic/gin"
	todorest "github.com/restlesswhy/todo-rest"
)

func (h *Handler) signUp(c *gin.Context) {
	var user todorest.User

	if err := c.BindJSON(&user); err != nil {
		
	}
}

func (h *Handler) signIn(c *gin.Context) {
	
}