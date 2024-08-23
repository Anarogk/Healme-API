package models

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"-"` // Hide password in JSON responses
	Specialization string `json:"specialization"`
	LicenseNumber  string `json:"license_number"`
}
