package interfaces

import "github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"

type AuthRepository interface {
	GetUserByEmail(email string) (entities.User, error)
	GetUserByID(id uint) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
}
