package services

import (
	"errors"

	"github.com/pteus/books-api/internal/models"
	"github.com/pteus/books-api/internal/repositories"
	"github.com/pteus/books-api/internal/utils"
)

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (a *authService) Register(username string, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
	}

	return a.userRepo.Create(user)
}

func (a *authService) Login(username string, password string) (string, error) {
	user, err := a.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	isValidPassword := utils.CheckPasswordHash(password, user.Password)

	if !isValidPassword {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
