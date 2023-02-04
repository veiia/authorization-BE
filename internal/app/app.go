package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/internal/handlers"
	"github.com/syth0le/authorization-BE/pkg/database"
	"os"
)

var db *sqlx.DB

func ConfigureApp() handlers.Routes {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	port := fmt.Sprintf(":%s", viper.GetString("port"))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	_, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error db connection: %s", err.Error())
	}

	r := handlers.NewRoutes(port)
	return r.CreateRoutes()
}
