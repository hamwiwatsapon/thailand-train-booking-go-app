package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique;not null;index"`
	PasswordHash string `json:"-"` // Store the hashed password
	Role         string `json:"role"`
}
