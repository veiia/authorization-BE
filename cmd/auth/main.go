package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/docs"
	"github.com/syth0le/authorization-BE/internal/app"
	"log"
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
	port := fmt.Sprintf(":%s", viper.GetString("port"))

	r := app.CreateRoutes()
	r.Run(port)
}
