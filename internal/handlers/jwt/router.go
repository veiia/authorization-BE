package jwt

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/jwt")

	r.GET("/get", getJWTTokenHandler)
	r.POST("/update", updateJWTTokenHandler)
	r.DELETE("/revoke", revokeJWTTokenHandler)
}

// GetJWTToken godoc
// @Summary      get jwt token
// @Description  get existing jwt token for user
// @Tags         jwt
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  jwtModel.JWTTokenResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/jwt/get [get]
func getJWTTokenHandler(c *gin.Context) {
	getJWTToken(c)
}

// UpdateJWTToken godoc
// @Summary      update jwt token
// @Description  update existing jwt token for user
// @Tags         jwt
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  jwtModel.JWTTokenResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/jwt/update [post]
func updateJWTTokenHandler(c *gin.Context) {
	updateJWTToken(c)
}

// revokeJWTTokenHandler godoc
// @Summary      revoke jwt token
// @Description  revoke existing jwt token for user
// @Tags         jwt
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  jwtModel.JWTTokenResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/jwt/revoke [delete]
func revokeJWTTokenHandler(c *gin.Context) {
	revokeJWTToken(c)
}
