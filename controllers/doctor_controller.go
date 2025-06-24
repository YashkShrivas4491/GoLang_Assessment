package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"makerble-crudapi/models"
	"makerble-crudapi/config" 
)

func GetAllPatientsForDoctor(c *gin.Context) {
	if c.GetString("role") != "doctor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func UpdatePatientForDoctor(c *gin.Context) {
	if c.GetString("role") != "doctor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	id := c.Param("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}
