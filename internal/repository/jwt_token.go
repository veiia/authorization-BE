package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/pkg/database"
)

type JwtTokenPostgres struct {
	db *sqlx.DB
}

func NewJwtTokenRepository(db *sqlx.DB) *JwtTokenPostgres {
	return &JwtTokenPostgres{db: db}
}

func (r *JwtTokenPostgres) Create(request jwtModel.JWTTokenCreateRequest) (jwtModel.JWTToken, error) {
	var token jwtModel.JWTToken
	query := fmt.Sprintf("CREATE * FROM %s WHERE token=%s", database.TokensTable, request.UserID)
	if err := r.db.Select(&token, query); err != nil {
		logrus.Error("Cannot create token", err.Error())
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}

func (r *JwtTokenPostgres) Revoke(request jwtModel.JWTTokenRequest) (int, error) {
	var token jwtModel.JWTToken
	query := fmt.Sprintf("UPDATE * FROM %s WHERE token=%s", database.TokensTable, request.Token)
	if err := r.db.Select(&token, query); err != nil {
		msg := fmt.Sprintf("TOKEN (%s) FOR USER (%s) NOT FOUND; ", request.Token, request.Username)
		logrus.Error(msg, err.Error())
		return 0, err
	}
	return token.UserId, nil
}

func (r *JwtTokenPostgres) Get(request jwtModel.JWTTokenRequest) (jwtModel.JWTToken, error) {
	var token jwtModel.JWTToken
	query := fmt.Sprintf("SELECT * FROM %s WHERE token=%s", database.TokensTable, request.Token)
	if err := r.db.Select(&token, query); err != nil {
		msg := fmt.Sprintf("TOKEN (%s) FOR USER (%s) NOT FOUND; ", request.Token, request.Username)
		logrus.Error(msg, err.Error())
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}

func (r *JwtTokenPostgres) Update(request jwtModel.JWTToken) (jwtModel.JWTToken, error) {
	var token jwtModel.JWTToken
	query := fmt.Sprintf("UPDATE * FROM %s WHERE token=%s", database.TokensTable, request.UserId)
	if err := r.db.Select(&token, query); err != nil {
		logrus.Error("Cannot create token", err.Error())
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}
