package core

import (
	"github.com/gin-gonic/gin"
	coreModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/internal/repository"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func signUp(c *gin.Context) {
	var input coreModel.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := repository.CreateUser(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := coreModel.UserIdResponse{Id: id}
	c.JSON(http.StatusOK, response)
}

func signIn(c *gin.Context) {
	var input coreModel.UserSignInRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func logOut(c *gin.Context) {
	var input coreModel.UserLogOutRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}
