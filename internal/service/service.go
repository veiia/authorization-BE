package service

import "github.com/syth0le/authorization-BE/internal/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
