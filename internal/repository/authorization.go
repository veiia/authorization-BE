package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/pkg/database"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (r *AuthorizationPostgres) SignUp(request userModel.SignUpRequest) (int, error) {
	//var user userModel.SignUpResponse

	return 0, nil
}

func (r *AuthorizationPostgres) SignIn(request userModel.SignInRequest) (int, error) {
	return 0, nil
}

func (r *AuthorizationPostgres) SignOut(request userModel.SignOutRequest) error {
	return nil
}

// create CreateUser, CreateTokenForUser, GetUser, GetTokenForUser, DeleteTokenForUser, UpdateTokenForUser

func (r *AuthorizationPostgres) GetUsers() ([]userModel.User, error) {
	var users []userModel.User
	query := fmt.Sprintf("SELECT * FROM %s", database.UsersTable)
	if err := r.db.Select(&users, query); err != nil {
		logrus.Error("USERS NOT FOUND; ", err.Error())
		return nil, err
	}
	return users, nil
}

//UPDATE customers
//SET first_name = 'Joseph'
//WHERE customer_id = 8000;

//UPDATE suppliers
//SET supplier_id   = 150,
//supplier_name = 'Apple',
//city          = 'Cupertino'
//state         = 'California'
//WHERE supplier_name = 'Google';
