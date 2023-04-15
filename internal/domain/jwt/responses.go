package jwtModel

type MessageResponse struct {
	Message string `json:"message"`
}

type JWTTokenAliveResponse struct {
	Token string `json:"token"`
	Alive bool   `json:"alive" default:"true"`
}

type JwtTokenRefreshResponse struct {
	Token string `json:"token"`
}
