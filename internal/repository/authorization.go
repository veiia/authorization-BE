package repository

import (
	"github.com/jmoiron/sqlx"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (r *AuthorizationPostgres) SignUp(request userModel.SignUpRequest) (userModel.SignUpResponse, error) {
	return userModel.SignUpResponse{}, nil
}

func (r *AuthorizationPostgres) SignIn(request userModel.SignInRequest) (userModel.SignInResponse, error) {
	return userModel.SignInResponse{}, nil
}

func (r *AuthorizationPostgres) SignOut(request userModel.SignOutRequest) (userModel.MessageResponse, error) {
	return userModel.MessageResponse{}, nil
}

func GetUsers() ([]userModel.User, error) {
	var users []userModel.User
	//query := fmt.Sprintf("SELECT * FROM %s", database.UsersTable)
	//if err := database.DB.Select(&users, query); err != nil {
	//	logrus.Error("USERS NOT FOUND; ", err.Error())
	//	return nil, err
	//}
	return users, nil
}
