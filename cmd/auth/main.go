package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/syth0le/authorization-BE/docs"
	"github.com/syth0le/authorization-BE/internal/app"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	docs.SwaggerInfo.Title = "authorization-BE API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "localhost:8008"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	port := app.ConfigureApp()
	r := app.CreateRoutes()

	err := r.Run(port)

	if err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
