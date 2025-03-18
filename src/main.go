package main

import (
	"fmt"
	"os"
	"rania-eskristal/src/commons/config"
	"rania-eskristal/src/infrastructures/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	level := os.Getenv("APP_LEVEL")
	if level == "" {
		level = "local"
	}

	conf := viper.New()
	config := config.New(conf, level)
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	fiber := http.New(config, logger)

	err := fiber.Listen(fmt.Sprintf("%v:%d", config.App.Host, config.App.Port))

	if err != nil {
		panic(fmt.Sprintf("error running the server: %v", err))
	}
}
