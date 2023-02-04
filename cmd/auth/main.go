package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syth0le/authorization-BE/docs"
	"github.com/syth0le/authorization-BE/internal/app"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	configureSwaggerInfo()

	r := app.ConfigureApp()

	err := r.Run(r.Port)
	if err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func configureSwaggerInfo() {
	docs.SwaggerInfo.Title = viper.GetString("title")
	docs.SwaggerInfo.Version = viper.GetString("version")
	docs.SwaggerInfo.BasePath = viper.GetString("base-path")
	docs.SwaggerInfo.Host = viper.GetString("host")
	docs.SwaggerInfo.Schemes = viper.GetStringSlice("schemes")
}
