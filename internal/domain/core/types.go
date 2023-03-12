package coreModel

type User struct {
	Id        int    `json:"-"`
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required" db:"encrypted_password"`
	Email     string `json:"email" binding:"required"`
	LastLogin string `json:"last_login" binding:"required"`
}
