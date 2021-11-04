package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/restlesswhy/todo-rest"
	"github.com/restlesswhy/todo-rest/handler"
	"github.com/restlesswhy/todo-rest/repository"
	"github.com/restlesswhy/todo-rest/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

  	if err := godotenv.Load(); err != nil {
    	logrus.Fatal("Error loading .env file")
  	}

	if err := InitConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func main() {
	srv := new(todorest.Server)

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		User: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	go func ()  {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			logrus.Fatalf("server is not run: %v", err)
		}
	}()

	logrus.Print("App started!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("App shutting down!")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	//TODO: full JWT, error support, makefile, docker compose
}

func InitConfig() error{
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig() // Find and read the config file
}