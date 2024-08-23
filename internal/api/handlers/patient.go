package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Anarogk/healthcare-api/internal/models"
)

type PatientHandler struct {
	DB *gorm.DB
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

func (h *PatientHandler) GetPatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) GetPatients(c *gin.Context) {
	var patients []models.Patient
	if err := h.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	h.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
