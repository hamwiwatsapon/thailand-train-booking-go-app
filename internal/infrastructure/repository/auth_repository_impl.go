package repository

import (
	"fmt"

	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/interfaces"
	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authRepository{db: db}
}

type authRepository struct {
	db *gorm.DB
}

// CreateUser implements interfaces.AuthRepository.
func (a *authRepository) CreateUser(user entities.User) (entities.User, error) {
	tx := a.db.Begin()
	if err := tx.Error; err != nil {
		return entities.User{}, err
	}

	if err := tx.Where("email = ?", user.Email).First(&entities.User{}).Error; err == nil {
		tx.Rollback()
		return entities.User{}, fmt.Errorf("email %s already exists", user.Email)
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return entities.User{}, err
	}

	return user, tx.Commit().Error
}

// DeleteUser implements interfaces.AuthRepository.
func (a *authRepository) DeleteUser(id uint) error {
	// Start a transaction
	tx := a.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// Check if the user exists and delete in one step
	if err := tx.Where("id = ?", id).Delete(&entities.User{}).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with id %d not found", id)
		}
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (a *authRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := a.db.Select("id", "email", "role").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return entities.User{}, err
	}
	return user, nil
}

// GetUserByID implements interfaces.AuthRepository.
func (a *authRepository) GetUserByID(id uint) (entities.User, error) {
	var user entities.User
	err := a.db.Select("id", "email", "role").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return entities.User{}, err
	}
	return user, nil
}

// GetUserByEmailWithDeleted implements interfaces.AuthRepository.
func (a *authRepository) GetUserByEmailWithDeleted(email string) (entities.User, error) {
	var user entities.User
	err := a.db.Unscoped().Select("id", "email", "role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, err
}

// GetUsers implements interfaces.AuthRepository.
func (a *authRepository) GetUsers(offset int, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	// Count total users
	if err := a.db.Model(&entities.User{}).Where("deleted_at IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated users
	err := a.db.Select("id", "email", "role").
		Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}

// GetUsersByRole implements interfaces.AuthRepository.
func (a *authRepository) GetUsersByRole(role string, offset int, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	// Count total users with the role
	if err := a.db.Model(&entities.User{}).
		Where("role = ? AND deleted_at IS NULL", role).
		Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users with role %s: %w", role, err)
	}

	// Fetch paginated users
	err := a.db.Select("id", "email", "role").
		Where("role = ? AND deleted_at IS NULL", role).
		Offset(offset).Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch users with role %s: %w", role, err)
	}

	return users, total, nil
}

// UpdateUser implements interfaces.AuthRepository.
func (a *authRepository) UpdateUser(user entities.User) (entities.User, error) {
	// Check if the user exists
	var existingUser entities.User
	if err := a.db.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, fmt.Errorf("user with id %d not found", user.ID)
		}
		return entities.User{}, fmt.Errorf("failed to fetch user with id %d: %w", user.ID, err)
	}

	// Update the user
	if err := a.db.Model(&existingUser).Updates(user).Error; err != nil {
		return entities.User{}, fmt.Errorf("failed to update user with id %d: %w", user.ID, err)
	}

	return existingUser, nil
}
