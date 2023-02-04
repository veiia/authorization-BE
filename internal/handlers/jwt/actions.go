package jwt

import (
	"github.com/gin-gonic/gin"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
)

func getJWTToken(c *gin.Context) {
	c.JSON(200, jwtModel.JWTTokenResponse{JWTToken: jwtModel.JWTToken{Token: "324234"}})
}

func updateJWTToken(c *gin.Context) {
	c.JSON(200, jwtModel.JWTTokenResponse{JWTToken: jwtModel.JWTToken{Token: "324234"}})
}

func revokeJWTToken(c *gin.Context) {
	c.JSON(200, jwtModel.JWTTokenResponse{JWTToken: jwtModel.JWTToken{Token: "324234"}})
}
