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
	sqlStr := "INSERT INTO ?(token, ttl, is_alive, user_id) values (?,?,?,?)"
	ret, err := r.db.Exec(sqlStr, database.TokensTable, "token", 15, true, request.UserID)
	if err != nil {
		logrus.Error("Cannot create token", err.Error())
		return jwtModel.JWTToken{}, err
	}
	insertedId, err := ret.LastInsertId()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", database.TokensTable, insertedId)
	if err := r.db.Select(&token, query); err != nil {
		msg := fmt.Sprintf("TOKEN (%s) NOT FOUND;", "token")
		logrus.Error(msg, err.Error())
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}

func (r *JwtTokenPostgres) Revoke(request jwtModel.JWTTokenRequest) (int, error) {
	sqlStr := "UPDATE ? SET is_alive=false WHERE token=? RETURNING *;"
	_, err := r.db.Exec(sqlStr, database.TokensTable, request.Token)
	if err != nil {
		msg := fmt.Sprintf("TOKEN (%s) FOR USER (%s) NOT FOUND; ", request.Token, request.Username)
		logrus.Error(msg, err.Error())
		return 0, err
	}
	var userId int
	return userId, nil
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
	query := fmt.Sprintf("UPDATE * FROM %s WHERE token=%d", database.TokensTable, request.UserId)
	if err := r.db.Select(&token, query); err != nil {
		logrus.Error("Cannot create token", err.Error())
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}
