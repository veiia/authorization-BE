package user

import (
	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup) {
	session := rg.Group("/user")

	session.GET("/", pongFunction)
}

func pongFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
