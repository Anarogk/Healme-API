package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Anarogk/healthcare-api/internal/models"
)

type DoctorHandler struct {
	DB *gorm.DB
}

func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := h.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func (h *DoctorHandler) GetDoctors(c *gin.Context) {
	var doctors []models.Doctor
	if err := h.DB.Find(&doctors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := h.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&doctor)
	c.JSON(http.StatusOK, doctor)
}

func (h *DoctorHandler) DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := h.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	h.DB.Delete(&doctor)
	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}
