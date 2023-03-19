package handler

import (
	"github.com/gin-gonic/gin"
	coreModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func SignUp(c *gin.Context, h *Handler) {
	var request coreModel.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	userId, err := h.Services.Authorization.CreateUser(request)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.Services.Authorization.GenerateToken(userId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response := coreModel.SignUpResponse{Username: request.Username, JwtToken: token.Token, IsAliveJwtToken: token.IsAlive}
	c.JSON(http.StatusOK, response)
}

func SignIn(c *gin.Context, h *Handler) {
	var request coreModel.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	userId, err := h.Services.Authorization.CheckSignIn(request)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response, err := h.Services.Authorization.GenerateToken(userId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Wrong password or sign-in. Try again")
	}
	c.JSON(http.StatusOK, response)
}

func SignOut(c *gin.Context, h *Handler) {
	var request coreModel.SignOutRequest
	if err := c.BindJSON(&request); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	err := h.Services.Authorization.SignOut(request)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.Status(http.StatusOK)
}

// for test
func GetUsers(c *gin.Context, h *Handler) {
	users, err := h.Services.Authorization.GetUsers()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, users)
}
