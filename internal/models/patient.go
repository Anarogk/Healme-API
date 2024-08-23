package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"-"` // Hide password in JSON responses
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
