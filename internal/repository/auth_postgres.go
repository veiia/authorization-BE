package repository

import (
	"fmt"
	"github.com/sirupsen/logrus"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/pkg/database"
)

func CreateUser(user userModel.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, username, encrypted_password) values ($1, $2, $3, $4) RETURNING id", database.UsersTable)
	row := database.DB.QueryRow(query, user.Email, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func GetUsers() ([]userModel.User, error) {
	var users []userModel.User
	query := fmt.Sprintf("SELECT * FROM %s", database.UsersTable)
	if err := database.DB.Select(&users, query); err != nil {
		logrus.Error("USERS NOT FOUND; ", err.Error())
		return nil, err
	}
	return users, nil
}
