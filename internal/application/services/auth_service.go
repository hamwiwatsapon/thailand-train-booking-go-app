package services

import (
	"errors"
	"strings"

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

	// Validate role (example: predefined roles)
	validRoles := map[string]bool{"admin": true, "user": true}
	if !validRoles[role] {
		return entities.User{}, errors.New("invalid role")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, errors.New("failed to hash password")
	}

	// Create the user
	user := entities.User{
		Email:    strings.ToLower(email),
		Password: string(hashedPassword),
		Role:     role,
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) LoginUser(email, password string) (string, string, error) {
	// Fetch the user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	token, refreshToken, err := GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, refreshToken, nil
}

func (s *AuthService) CheckUserExist(email string) (string, error) {
	// Check if the user exists in the database
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.ID == 0 {
		return "", errors.New("user not found")
	}

	otp, ref, err := GenerateOTP(email)
	if err != nil {
		return "", errors.New("failed to generate OTP")
	}

	err = SendOTPEmail(email, ref, otp)
	if err != nil {
		return "", errors.New("failed to send OTP email")
	}

	return ref, nil
}

func (s *AuthService) OTPLogin(email, otp string) (string, string, error) {
	// Fetch the user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email or OTP")
	}

	// Validate the OTP (this is just a placeholder, implement your own OTP validation logic)
	validate, err := ValidateOTP(email, otp)
	if err != nil || !validate {
		return "", "", errors.New("invalid email or OTP")
	}

	token, responseToken, err := GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, responseToken, nil
}

func (s *AuthService) GetNewToken(refreshToken string) (string, string, error) {
	// Validate the token and extract the user ID
	token, refreshToken, err := RefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid token")
	}

	return token, refreshToken, nil
}
