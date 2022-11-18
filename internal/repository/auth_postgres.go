package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	userModel "github.com/syth0le/authorization-BE/internal/domain/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user userModel.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, username, encrypted_password) values ($1, $2, $3, $4) RETURNING id")
	row := r.db.QueryRow(query, user.Email, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
