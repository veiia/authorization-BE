package jwt

import (
	"github.com/gin-gonic/gin"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/internal/handler"
	"github.com/syth0le/authorization-BE/pkg/utils"
	"net/http"
)

func RefreshJWTToken(c *gin.Context, h *handler.Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, requestBody)
}

func RevokeJWTToken(c *gin.Context, h *handler.Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusNoContent, jwtModel.MessageResponse{Message: "JWT Token was revoked"})
}

func IsJWTTokenAlive(c *gin.Context, h *handler.Handler) {
	var requestBody jwtModel.JWTTokenRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, requestBody)
}
