package services_test

import (
	"errors"
	"testing"

	"github.com/hamwiwatsapon/train-booking-go/internal/application/services"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mock repository
type MockAuthRepository struct {
	mock.Mock
}

// CreateUser implements interfaces.AuthRepository.
func (m *MockAuthRepository) CreateUser(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

// DeleteUser implements interfaces.AuthRepository.
func (m *MockAuthRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(1)
}

// GetUserByEmail implements interfaces.AuthRepository.
func (m *MockAuthRepository) GetUserByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

// GetUserByEmailWithDeleted implements interfaces.AuthRepository.
func (m *MockAuthRepository) GetUserByEmailWithDeleted(email string) (entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

// GetUserByID implements interfaces.AuthRepository.
func (m *MockAuthRepository) GetUserByID(id uint) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

// GetUsers implements interfaces.AuthRepository.
func (m *MockAuthRepository) GetUsers(offset int, limit int) ([]entities.User, int64, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]entities.User), args.Get(1).(int64), args.Error(2)
}

// GetUsersByRole implements interfaces.AuthRepository.
func (m *MockAuthRepository) GetUsersByRole(role string, offset int, limit int) ([]entities.User, int64, error) {
	args := m.Called(role, offset, limit)
	return args.Get(0).([]entities.User), args.Get(1).(int64), args.Error(2)
}

// UpdateUser implements interfaces.AuthRepository.
func (m *MockAuthRepository) UpdateUser(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func TestRegisterUserSuccess(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := services.NewAuthService(mockRepo)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(entities.User{}, errors.New("not found"))
	mockRepo.On("CreateUser", mock.Anything).Return(entities.User{Email: "test@example.com"}, nil)

	user, err := service.RegisterUser("test@example.com", "password123", "user")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestRegisterUserUserAlreadyExists(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := services.NewAuthService(mockRepo)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(entities.User{Email: "test@example.com"}, nil)

	user, err := service.RegisterUser("test@example.com", "password123", "user")
	assert.Error(t, err)
	assert.Equal(t, "email already exists", err.Error())
	assert.Empty(t, user)
}

func TestLoginSuccess(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := services.NewAuthService(mockRepo)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mockRepo.On("GetUserByEmail", "test@example.com").Return(entities.User{Email: "test@example.com", PasswordHash: string(hashedPassword)}, nil)

	user, err := service.LoginUser("test@example.com", "password123")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestLoginFail(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	service := services.NewAuthService(mockRepo)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(entities.User{}, errors.New("not found"))

	user, err := service.LoginUser("test@example.com", "wrongpassword")
	assert.Error(t, err)
	assert.Equal(t, "invalid email or password", err.Error())
	assert.Empty(t, user)
}
