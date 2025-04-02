package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"not null; default: user"`
}
