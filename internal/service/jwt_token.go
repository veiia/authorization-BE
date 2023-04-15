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

func (s *JwtTokenService) Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error) {
	token, err := s.repo.Get(request)
	if err != nil {
		return jwtModel.JWTTokenAliveResponse{}, err
	}
	// TODO logic для определения жив ли токен
	// сделать IsAlive приходящий из репы пропертей что вычисляется с помощью aggregate() и в базе хранить ее не стоит
	//if token.Ttl + Time > now()
	//	token.IsAlive = False
	return jwtModel.JWTTokenAliveResponse{Token: token.Token, Alive: token.IsAlive}, nil
}

func (s *JwtTokenService) Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error) {
	userId, err := s.repo.Revoke(request)
	if err != nil {
		return jwtModel.JwtTokenRefreshResponse{}, err
	}
	token, err := s.repo.Create(jwtModel.JWTTokenCreateRequest{UserID: userId})
	if err != nil {
		return jwtModel.JwtTokenRefreshResponse{}, err
	}
	return jwtModel.JwtTokenRefreshResponse{Token: token.Token}, nil
}

func (s *JwtTokenService) Revoke(request jwtModel.JWTTokenRequest) (jwtModel.MessageResponse, error) {
	_, err := s.repo.Revoke(request)
	if err != nil {
		return jwtModel.MessageResponse{Message: "JWT Token was not revoked. Error was occurred."}, nil
	}
	return jwtModel.MessageResponse{Message: "JWT Token was revoked"}, nil
}
