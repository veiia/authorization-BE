package service

import (
	"crypto/sha1"
	"fmt"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	"github.com/syth0le/authorization-BE/internal/repository"
	"time"
)

const (
	salt       = "dasfhbiy3h1hibfcvi2b4h2bfhb2h"
	signingKey = "qrkjk#4#%adshb34jb1b#4n#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(user userModel.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return 1, nil
}

func (s *AuthorizationService) GenerateToken(username, password string) (string, error) {
	return "", nil
}

func (s *AuthorizationService) ParseToken(token string) (int, error) {
	return 1, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
