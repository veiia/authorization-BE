package repository

import (
	"github.com/jmoiron/sqlx"
	userModel "github.com/syth0le/authorization-BE/internal/domain/user"
)

type Authorization interface {
	CreateUser(user userModel.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
