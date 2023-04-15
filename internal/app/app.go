package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/internal/handler"
	"github.com/syth0le/authorization-BE/internal/repository"
	"github.com/syth0le/authorization-BE/internal/server"
	"github.com/syth0le/authorization-BE/internal/service"
	"github.com/syth0le/authorization-BE/pkg/database"
	"os"
	"os/signal"
)

type App struct {
	db       *sqlx.DB
	handlers *handler.Handler
}

func NewApp() *App {
	db := initDb()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	return &App{
		db,
		handlers,
	}
}

func (a *App) Run(port string) error {
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
		return err
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
		return err
	}

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), a.handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
		return err
	}

	if err := a.db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
		return err
	}
	return nil
}

func initDb() *sqlx.DB {
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
	return db
}
