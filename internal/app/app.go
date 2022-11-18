package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syth0le/authorization-BE/config"
	"github.com/syth0le/authorization-BE/internal/handlers/oauth"
	"github.com/syth0le/authorization-BE/internal/handlers/session"
	"github.com/syth0le/authorization-BE/internal/handlers/user"
	"github.com/syth0le/authorization-BE/pkg/database"
	"os"
)

type Routes struct {
	router *gin.Engine
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}

func CreateRoutes() Routes {
	r := Routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	session.AddRoute(v1)
	user.AddRoute(v1)
	oauth.AddRoute(v1)

	return r
}

func ConfigureApp() string {
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
	//repos := repository.NewRepository(db)
	//service := service.NewService(repos)
	//_ = handlers.NewHandler(service)
	return port
}
