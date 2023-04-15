package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/syth0le/authorization-BE/internal/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1/auth")

	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	AddCoreRouter(v1, h)
	AddJwtRouter(v1, h)

	return router
}
