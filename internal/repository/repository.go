package repository

import (
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
)

type Authorization interface {
	CreateUser(user userModel.User) (int, error)
}

type Repository struct {
	Authorization
}
