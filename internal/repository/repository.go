package repository

import (
	"github.com/jmoiron/sqlx"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
)

type Authorization interface {
	SignUp(request userModel.SignUpRequest) (int, error)
	SignIn(request userModel.SignInRequest) (int, error)
	SignOut(request userModel.SignOutRequest) error
	GetUsers() ([]userModel.User, error)
}

type JwtToken interface {
	Create(request jwtModel.JWTTokenCreateRequest) (jwtModel.JWTToken, error)
	Revoke(request jwtModel.JWTTokenRequest) (int, error)
	Get(request jwtModel.JWTTokenRequest) (jwtModel.JWTToken, error)
	Update(request jwtModel.JWTToken) (jwtModel.JWTToken, error)
}

type Repository struct {
	Authorization
	JwtToken
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepository(db),
		JwtToken:      NewJwtTokenRepository(db),
	}
}
