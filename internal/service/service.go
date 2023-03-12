package service

import (
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/internal/repository"
)

type Authorization interface {
	CreateUser(user userModel.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type JwtToken interface {
	Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error)
	Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error)
	Revoke(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error)
}

type Service struct {
	Authorization
	JwtToken
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		JwtToken:      NewJwtTokenService(repos.JwtToken),
	}
}
