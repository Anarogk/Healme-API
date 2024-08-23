package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Anarogk/healthcare-api/internal/models"
)

type AppointmentHandler struct {
	DB *gorm.DB
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the doctor is available at the requested time
	var conflictingAppointment models.Appointment
	if err := h.DB.Where("doctor_id = ? AND date_time = ?", appointment.DoctorID, appointment.DateTime).First(&conflictingAppointment).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Doctor is not available at the requested time"})
		return
	}

	// Check if the patient has any conflicting appointments
	if err := h.DB.Where("patient_id = ? AND date_time = ?", appointment.PatientID, appointment.DateTime).First(&conflictingAppointment).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Patient already has an appointment at the requested time"})
		return
	}

	// Set the appointment status to "Scheduled"
	appointment.Status = "Scheduled"

	if err := h.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	query := h.DB.Preload("Patient").Preload("Doctor")

	// Filter appointments by date range
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if startDate != "" && endDate != "" {
		query = query.Where("date_time BETWEEN ? AND ?", startDate, endDate)
	}

	// Filter appointments by doctor or patient
	doctorID := c.Query("doctor_id")
	patientID := c.Query("patient_id")
	if doctorID != "" {
		query = query.Where("doctor_id = ?", doctorID)
	}
	if patientID != "" {
		query = query.Where("patient_id = ?", patientID)
	}

	if err := query.Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := h.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	var updateData struct {
		DateTime    string `json:"date_time"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the new time is available for the doctor
	if updateData.DateTime != "" && updateData.DateTime != appointment.DateTime {
		var conflictingAppointment models.Appointment
		if err := h.DB.Where("doctor_id = ? AND date_time = ? AND id != ?", appointment.DoctorID, updateData.DateTime, appointment.ID).First(&conflictingAppointment).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Doctor is not available at the requested time"})
			return
		}
		appointment.DateTime = updateData.DateTime
	}

	if updateData.Description != "" {
		appointment.Description = updateData.Description
	}

	if updateData.Status != "" {
		appointment.Status = updateData.Status
	}

	h.DB.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment
	if err := h.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	h.DB.Delete(&appointment)
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
