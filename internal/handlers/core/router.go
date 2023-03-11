package core

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/")

	r.POST("/sign-up", signUpHandler)
	r.POST("/sign-in", signInHandler)
	r.POST("/sign-out", signOutHandler)
	r.GET("/users", getUsersHandler)
}

// SignUp godoc
// @Summary      Sign up core
// @Description  route for signing up to veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.User  true  "User Info"
//
//	@Success      200         {object}  coreModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-up [post]
func signUpHandler(c *gin.Context) {
	signUp(c)
}

// SignIn godoc
// @Summary      Sign in core
// @Description  route for signing in to veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.UserSignInRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-in [post]
func signInHandler(c *gin.Context) {
	signIn(c)
}

// signOut godoc
// @Summary      Sign Out core
// @Description  route for signging out from veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.UsersignOutRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/auth/sign-out [post]
func signOutHandler(c *gin.Context) {
	signOut(c)
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
func getUsersHandler(c *gin.Context) {
	getUsers(c)
}
