package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
