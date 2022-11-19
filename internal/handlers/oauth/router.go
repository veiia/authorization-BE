package oauth

import "github.com/gin-gonic/gin"

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/oauth")
	_ = r
}
