package service

import (
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/internal/repository"
)

type JwtTokenService struct {
	repo repository.JwtToken
}

func NewJwtTokenService(repo repository.JwtToken) *JwtTokenService {
	return &JwtTokenService{repo: repo}
}

func (r *JwtTokenService) Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error) {
	return jwtModel.JWTTokenAliveResponse{}, nil
}

func (r *JwtTokenService) Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error) {
	return jwtModel.JwtTokenRefreshResponse{}, nil
}

func (r *JwtTokenService) Revoke(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error) {
	return jwtModel.JwtTokenRefreshResponse{}, nil
}
