package jwt

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/jwt")

	r.POST("/refresh", refreshJWTTokenHandler)
	r.DELETE("/revoke", revokeJWTTokenHandler)
	r.POST("/alive", isJWTTokenAliveHandler)
}

// UpdateJWTToken godoc
// @Summary      refresh jwt token
// @Description  update existing jwt token for user
// @Tags         jwt
// @Accept       json
// @Produce      json
// @Param request body jwtModel.JWTTokenRequest true "auth params"
//
//	@Success      200         {object}  jwtModel.JwtTokenRefreshResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/jwt/refresh [post]
func refreshJWTTokenHandler(c *gin.Context) {
	refreshJWTToken(c)
}

// revokeJWTTokenHandler godoc
// @Summary      revoke jwt token
// @Description  revoke existing jwt token for user
// @Tags         jwt
// @Accept       json
// @Produce      json
// @Param request body jwtModel.JWTTokenRequest true "auth params"
//
//	@Success      200         {object}  jwtModel.MessageResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/jwt/revoke [delete]
func revokeJWTTokenHandler(c *gin.Context) {
	revokeJWTToken(c)
}

// isJWTTokenAliveHandler godoc
// @Summary      check jwt token status
// @Description  check if jwt token alive return json response, else 401 Unauthorized
// @Tags         jwt
// @Accept       json
// @Produce      json
// @Param request body jwtModel.JWTTokenRequest true "auth params"
//
//	@Success      200         {object}  jwtModel.JWTTokenAliveResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      401         {string}  string  "Unauthorized. JWT Expired"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/jwt/alive [post]
func isJWTTokenAliveHandler(c *gin.Context) {
	isJWTTokenAlive(c)
}
