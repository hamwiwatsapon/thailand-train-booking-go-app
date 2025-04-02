package interfaces

import "github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"

type AuthRepository interface {
	// Basic CRUD operations
	GetUserByEmail(email string) (entities.User, error)
	GetUserByID(id uint) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)

	// Additional operations
	DeleteUser(id uint) error
	GetUsers(offset, limit int) ([]entities.User, int64, error)
	GetUsersByRole(role string, offset, limit int) ([]entities.User, int64, error)
	GetUserByEmailWithDeleted(email string) (entities.User, error)
}
