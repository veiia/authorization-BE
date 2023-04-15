package handler

import (
	"github.com/gin-gonic/gin"
)

func AddCoreRouter(rg *gin.RouterGroup, h *Handler) {
	r := rg.Group("/")

	r.POST("/sign-up", h.signUpHandler)
	r.POST("/sign-in", h.signInHandler)
	r.POST("/sign-out", h.signOutHandler)
	r.GET("/users", h.getUsersHandler)
}

// SignUp godoc
// @Summary      Sign up core
// @Description  route for signing up to veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.SignUpRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.SignUpResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-up [post]
func (h *Handler) signUpHandler(c *gin.Context) {
	SignUp(c, h)
}

// SignIn godoc
// @Summary      Sign in core
// @Description  route for signing in to veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.SignInRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.SignInResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-in [post]
func (h *Handler) signInHandler(c *gin.Context) {
	SignIn(c, h)
}

// signOut godoc
// @Summary      Sign Out core
// @Description  route for signing out from veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.SignOutRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.MessageResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-out [post]
func (h *Handler) signOutHandler(c *gin.Context) {
	SignOut(c, h)
}

// TEST godoc
// @Summary      TEST route
// @Description  route for TEST
// @Tags         TEST
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  coreModel.User
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/users [get]
func (h *Handler) getUsersHandler(c *gin.Context) {
	GetUsers(c, h)
}
