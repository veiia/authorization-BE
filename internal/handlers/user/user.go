package user

import (
	"github.com/gin-gonic/gin"
	userModel "github.com/syth0le/authorization-BE/internal/domain/user"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func AddUserRoute(rg *gin.RouterGroup) {
	session := rg.Group("/user")

	session.POST("/sign-up", signUp)
	session.POST("/sign-in", signIn)
}

func signUp(c *gin.Context) {
	var input userModel.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func signIn(c *gin.Context) {
	var input userModel.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}
