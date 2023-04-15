package service

import (
	"crypto/sha1"
	"fmt"
	userModel "github.com/syth0le/authorization-BE/internal/domain/core"
	jwtModel "github.com/syth0le/authorization-BE/internal/domain/jwt"
	"github.com/syth0le/authorization-BE/internal/repository"
	"time"
)

const (
	salt       = "dasfhbiy3h1hibfcvi2b4h2bfhb2h"
	signingKey = "qrkjk#4#%adshb34jb1b#4n#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthorizationService struct {
	repo      repository.Authorization
	repoToken repository.JwtToken
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(request userModel.SignUpRequest) (int, error) {
	request.Password = generatePasswordHash(request.Password)
	userId, err := s.repo.SignUp(request)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *AuthorizationService) GenerateToken(userId int) (jwtModel.JWTToken, error) {
	token, err := s.repoToken.Create(jwtModel.JWTTokenCreateRequest{UserID: userId})
	if err != nil {
		return jwtModel.JWTToken{}, err
	}
	return token, nil
}

func (s *AuthorizationService) SignOut(request userModel.SignOutRequest) error {
	return s.repo.SignOut(request)
}

func (s *AuthorizationService) CheckSignIn(request userModel.SignInRequest) (int, error) {
	request.Password = generatePasswordHash(request.Password)
	userId, err := s.repo.SignIn(request)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *AuthorizationService) GetUsers() ([]userModel.User, error) {
	res, err := s.repo.GetUsers()
	fmt.Printf("%v", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
