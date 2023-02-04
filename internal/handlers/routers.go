package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syth0le/authorization-BE/internal/handlers/core"
	"github.com/syth0le/authorization-BE/internal/handlers/jwt"
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

	core.AddRouter(v1)
	jwt.AddRouter(v1)
	return r
}
