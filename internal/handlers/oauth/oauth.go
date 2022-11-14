package oauth

import (
	"github.com/gin-gonic/gin"
)

func AddOAuthRoute(rg *gin.RouterGroup) {
	session := rg.Group("/oauth")

	session.GET("/", pongFunction)
}

func pongFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
