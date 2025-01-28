package main

import (
	"github.com/didsqq/ZeroAgency_Test/internal/handler"
	"github.com/didsqq/ZeroAgency_Test/internal/repository"
	"github.com/didsqq/ZeroAgency_Test/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
		return
	}

	reformDB := reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logrus.Debugf))

	repos := NewRepository(reformDB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := handlers.InitRoutes()

	app.Listen()
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
