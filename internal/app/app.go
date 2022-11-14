package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syth0le/authorization-BE/internal/handlers/oauth"
	"github.com/syth0le/authorization-BE/internal/handlers/session"
	"github.com/syth0le/authorization-BE/internal/handlers/user"
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
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	session.AddSessionRoute(v1)
	user.AddUserRoute(v1)
	oauth.AddOAuthRoute(v1)

	return r
}
