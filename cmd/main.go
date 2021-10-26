package main

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/handler"
	"github.com/restlesswhy/todo-rest/repository"
	"github.com/restlesswhy/todo-rest/service"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := new(todorest.Server)

	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("server is not run: %v", err)
	}
}