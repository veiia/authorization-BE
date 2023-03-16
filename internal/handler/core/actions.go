package core

import (
	"github.com/gin-gonic/gin"
	coreModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/internal/handler"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func SignUp(c *gin.Context, h *handler.Handler) {
	var request coreModel.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	userId, err := h.Services.Authorization.CreateUser(request.(coreModel.User))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	token, err := h.Services.Authorization.GenerateToken(request.Username, request.Password)
	if err != nil {
		return
	}
	response := coreModel.SignUpResponse{Username: request.Username, JwtToken: token, IsAliveJwtToken: token.IsAlive}
	c.JSON(http.StatusOK, response)
}

func SignIn(c *gin.Context, h *handler.Handler) {
	var input coreModel.SignInRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response, err := h.Services.Authorization.GenerateToken(input)
	//if err != nil {
	//	utils.NewErrorResponse(c, http.StatusBadRequest, "Wrong password or sign-in. Try again")
	//}
	//c.JSON(http.StatusOK, response) // TODO: return JWTToken
}

func SignOut(c *gin.Context, h *handler.Handler) {
	var input coreModel.SignOutRequest
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

// for test
func GetUsers(c *gin.Context, h *handler.Handler) {
	//users, err := repository.GetUsers()
	//if err != nil {
	//	utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	//}
	//c.JSON(http.StatusOK, users)
}
