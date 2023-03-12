package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syth0le/authorization-BE/internal/handler/core"
	"github.com/syth0le/authorization-BE/internal/handler/jwt"
	"github.com/syth0le/authorization-BE/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1/auth")

	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	core.AddRouter(v1)
	jwt.AddRouter(v1)

	return router
}
