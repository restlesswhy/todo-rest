package main

import (
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/handler"
	"github.com/restlesswhy/todo-rest/repository"
	"github.com/restlesswhy/todo-rest/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	srv := new(todorest.Server)

	srv.InitConfig()

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		User: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("server is not run: %v", err)
	}
}