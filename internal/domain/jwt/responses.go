package jwtModel

type MessageResponse struct {
	Message string `json:"message"`
}

type JWTTokenAliveResponse struct {
	Alive bool `json:"alive" default:"true"`
}

type JwtTokenRefreshResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
