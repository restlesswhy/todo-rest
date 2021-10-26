package main

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := new(todorest.Server)
	
	handler := handler.NewHandler()

	if err := srv.Run("8080", ); err != nil {
		logrus.Fatalf("server is not run: %v", err)
	}
}