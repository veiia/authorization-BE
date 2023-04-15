package handler

import (
	"github.com/gin-gonic/gin"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func RefreshJWTToken(c *gin.Context, h *Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	newToken, err := h.Services.Refresh(requestBody)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, newToken)
}

func RevokeJWTToken(c *gin.Context, h *Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	isRevoked, err := h.Services.Revoke(requestBody)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusNoContent, isRevoked)
}

func IsJWTTokenAlive(c *gin.Context, h *Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	alive, err := h.Services.Alive(requestBody)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, alive)
}
