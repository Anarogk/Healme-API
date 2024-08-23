package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID   uint
	Patient     Patient
	DoctorID    uint
	Doctor      Doctor
	DateTime    string
	Description string
	Status      string // e.g., "Scheduled", "Completed", "Cancelled"
}
