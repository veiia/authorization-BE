package service

import (
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/internal/repository"
)

type Authorization interface {
	CreateUser(user userModel.SignUpRequest) (int, error)
	GenerateToken(userId int) (jwtModel.JWTToken, error)
	SignOut(request userModel.SignOutRequest) error
	CheckSignIn(request userModel.SignInRequest) (int, error)
	GetUsers() ([]userModel.User, error)
}

type JwtToken interface {
	Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error)
	Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error)
	Revoke(request jwtModel.JWTTokenRequest) (jwtModel.MessageResponse, error)
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
