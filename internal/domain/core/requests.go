package coreModel

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignOutRequest struct {
	Username string `json:"username" binding:"required"`
	JwtToken string `json:"jwt_token" binding:"required"`
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
