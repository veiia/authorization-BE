package handler

import (
	"github.com/gin-gonic/gin"
)

func AddJwtRouter(rg *gin.RouterGroup, h *Handler) {
	r := rg.Group("/jwt")

	r.POST("/refresh", h.refreshJWTTokenHandler)
	r.DELETE("/revoke", h.revokeJWTTokenHandler)
	r.POST("/alive", h.isJWTTokenAliveHandler)
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
func (h *Handler) refreshJWTTokenHandler(c *gin.Context) {
	RefreshJWTToken(c, h)
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
func (h *Handler) revokeJWTTokenHandler(c *gin.Context) {
	RevokeJWTToken(c, h)
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
func (h *Handler) isJWTTokenAliveHandler(c *gin.Context) {
	IsJWTTokenAlive(c, h)
}
