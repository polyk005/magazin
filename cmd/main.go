package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/polyk005/magazin/pkg/server"
	"github.com/polyk005/magazin/pkg/handler"
	"github.com/polyk005/magazin/pkg/repository"
	"github.com/polyk005/magazin/pkg/service"
)

func main() {
	fmt.Println("Starting server...")
	logrus.SetFormatter(new(logrus.JSONFormatter))
	
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	srv := new(magazin.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}