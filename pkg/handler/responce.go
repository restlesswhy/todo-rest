package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StatusResponce struct {
	Status string `json:"status"`
}

type ErrorResponce struct {
	Message string `json:"message"`
}

func NewErrorResponce(c *gin.Context, code int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(code, ErrorResponce{message})
}


