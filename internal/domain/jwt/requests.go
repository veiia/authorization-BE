package jwtModel

type JWTTokenRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type JWTTokenCreateRequest struct {
	UserID int `json:"user_id" binding:"required"`
}
