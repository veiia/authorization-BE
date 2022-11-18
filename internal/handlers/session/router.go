package session

import "github.com/gin-gonic/gin"

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/session")
	_ = r
}
