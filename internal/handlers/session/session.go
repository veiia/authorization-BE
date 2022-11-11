package session

import (
	"github.com/gin-gonic/gin"
)

func AddSessionRoute(rg *gin.RouterGroup) {
	session := rg.Group("/session")

	session.GET("/", pongFunction)
}

func pongFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
