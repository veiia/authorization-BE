package repository

import (
	"github.com/jmoiron/sqlx"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
)

type JwtTokenPostgres struct {
	db *sqlx.DB
}

func NewJwtTokenRepository(db *sqlx.DB) *JwtTokenPostgres {
	return &JwtTokenPostgres{db: db}
}

func (r *JwtTokenPostgres) Alive(request jwtModel.JWTTokenRequest) (jwtModel.JWTTokenAliveResponse, error) {
	return jwtModel.JWTTokenAliveResponse{}, nil
}

func (r *JwtTokenPostgres) Refresh(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error) {
	return jwtModel.JwtTokenRefreshResponse{}, nil
}

func (r *JwtTokenPostgres) Revoke(request jwtModel.JWTTokenRequest) (jwtModel.JwtTokenRefreshResponse, error) {
	return jwtModel.JwtTokenRefreshResponse{}, nil
}
