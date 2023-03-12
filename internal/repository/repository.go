package repository

import (
	"github.com/jmoiron/sqlx"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
)

type Authorization interface {
	SignUp(request userModel.SignUpRequest) (userModel.SignUpResponse, error)
	SignIn(request userModel.SignInRequest) (userModel.SignInResponse, error)
	SignOut(request userModel.SignOutRequest) (userModel.MessageResponse, error)
}

type JwtToken interface {
	Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error)
	Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error)
	Revoke(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error)
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
