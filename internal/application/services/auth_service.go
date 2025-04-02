package services

import (
	"errors"

	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo interfaces.AuthRepository
}

func NewAuthService(repo interfaces.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(email, password, role string) (entities.User, error) {
	// Check if the email already exists
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		return entities.User{}, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}

	// Create the user
	user := entities.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) LoginUser(email, password string) (entities.User, error) {
	// Fetch the user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return entities.User{}, errors.New("invalid email or password")
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return entities.User{}, errors.New("invalid email or password")
	}

	return user, nil
}
