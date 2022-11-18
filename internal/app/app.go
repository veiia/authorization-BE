package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/internal/handlers"
	"github.com/syth0le/authorization-BE/internal/repository"
	"github.com/syth0le/authorization-BE/internal/service"
	"github.com/syth0le/authorization-BE/pkg/database"
	"os"
)

func ConfigureApp() handlers.Routes {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	port := fmt.Sprintf(":%s", viper.GetString("port"))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := database.NewPostgresDB(database.Config{
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
	_ = db
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//_ := handlers.NewHandler(services)
	_ = services
	r := handlers.NewRoutes(port)
	return r.CreateRoutes()
}
