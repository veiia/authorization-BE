package user

import (
	"github.com/gin-gonic/gin"
	"github.com/syth0le/authorization-BE/internal/repository"
	"net/http"
)

func AddRouter(rg *gin.RouterGroup) {
	r := rg.Group("/user")

	r.POST("/sign-up", signUpHandler)
	r.POST("/sign-in", signInHandler)
	r.GET("/list", GetUsersHandler)
}

// SignUp godoc
// @Summary      Sign up user
// @Description  sign up
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        message  body      userModel.User  true  "User Info"
//
//	@Success      200         {object}  userModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /user/sign-up [post]
func signUpHandler(c *gin.Context) {
	signUp(c)
}

// SignIn godoc
// @Summary      Sign in user
// @Description  sign in
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        message  body      userModel.UserSignIn  true  "User Info"
//
//	@Success      200         {object}  userModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /user/sign-in [post]
func signInHandler(c *gin.Context) {
	signIn(c)
}

// GetUsersHandler godoc
// @Summary      Get users
// @Description  get list of users
// @Tags         user
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  userModel.UserIdResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /user/list [get]
func GetUsersHandler(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, users)
}
