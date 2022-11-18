package handlers

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
	Port   string
}

func NewRoutes(port string) *Routes {
	return &Routes{
		router: gin.Default(),
		Port:   port,
	}
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}

func (r Routes) CreateRoutes() Routes {

	v1 := r.router.Group("/v1")
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	session.AddRouter(v1)
	user.AddRouter(v1)
	oauth.AddRouter(v1)

	return r
}
