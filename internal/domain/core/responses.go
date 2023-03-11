package coreModel

type UserIdResponse struct {
	Id int `json:"id"`
}

type SignInResponse struct {
	Username        int    `json:"username"`
	JwtToken        string `json:"jwt_token"`
	IsAliveJwtToken bool   `json:"is_alive_jwt_token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SignUpResponse struct {
	Username        int    `json:"username"`
	JwtToken        string `json:"jwt_token"`
	IsAliveJwtToken bool   `json:"is_alive_jwt_token"`
}
