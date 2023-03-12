package jwtModel

type JWTToken struct {
	Id      int    `json:"-"`
	Token   string `json:"token" binding:"required"`
	Ttl     string `json:"ttl" binding:"required"`
	IsAlive bool   `json:"is_alive" binding:"required" db:"is_alive"`
	UserId  int    `json:"user_id" binding:"required" db:"user_id"`
}
