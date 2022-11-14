package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/docs"
	"github.com/syth0le/authorization-BE/internal/app"
	"github.com/syth0le/authorization-BE/internal/repository"
	"github.com/syth0le/authorization-BE/internal/service"
	"github.com/syth0le/authorization-BE/pkg/database"
	"log"
	"os"
)

func main() {
	docs.SwaggerInfo.Title = "authorization-BE API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "localhost:8008"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err := config.InitConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	port := fmt.Sprintf(":%s", viper.GetString("port"))
	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Printf("DLSAFLADKFLADFLJADFJLDAJLFN")
		log.Fatalf("error db connection: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	_ = service.NewService(repos)
	//handlers := handlers.NewHandler

	r := app.CreateRoutes()
	err = r.Run(port)
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
