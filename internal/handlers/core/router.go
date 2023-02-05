package core

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/")

	r.POST("/sign-up", signUpHandler)
	r.POST("/sign-in", signInHandler)
	r.POST("/log-out", logOutHandler)
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
// @Router       /v1/sign-up [post]
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
// @Router       /v1/sign-in [post]
func signInHandler(c *gin.Context) {
	signIn(c)
}

// LogOut godoc
// @Summary      Log Out core
// @Description  route for logging out from veiia system
// @Tags         core
// @Accept       json
// @Produce      json
// @Param        message  body      coreModel.UserLogOutRequest  true  "User Info"
//
//	@Success      200         {object}  coreModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /v1/log-out [post]
func logOutHandler(c *gin.Context) {
	logOut(c)
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
// @Router       /v1/users [get]
func getUsersHandler(c *gin.Context) {
	getUsers(c)
}
