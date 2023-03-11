package core

import (
	"github.com/gin-gonic/gin"
	coreModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/internal/repository"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func signUp(c *gin.Context) {
	var input coreModel.SignUpRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := repository.CreateUser(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := coreModel.UserIdResponse{Id: id} // TODO: return JWTToken
	c.JSON(http.StatusOK, response)
}

func signIn(c *gin.Context) {
	var input coreModel.SignInRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response, err := repository.CheckUserSignIn(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Wrong password or sign-in. Try again")
	}
	c.JSON(http.StatusOK, response) // TODO: return JWTToken
}

func signOut(c *gin.Context) {
	var input coreModel.SignOutRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

// for test
func getUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, users)
}
