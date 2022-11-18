package user

import (
	"github.com/gin-gonic/gin"
	userModel "github.com/syth0le/authorization-BE/internal/domain/user"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func AddRoute(rg *gin.RouterGroup) {
	session := rg.Group("/user")

	session.POST("/sign-up", signUp)
	session.POST("/sign-in", signIn)
}

func signUp(c *gin.Context) {
	var input userModel.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	//id, err := authCreateUser(input)
	id := 1
	//if err != nil {
	//	utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func signIn(c *gin.Context) {
	var input userModel.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}
