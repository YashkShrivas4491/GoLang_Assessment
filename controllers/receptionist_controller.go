package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"makerble-crudapi/models"
	"makerble-crudapi/config"
)

func CreatePatient(c *gin.Context) {
	if c.GetString("role") != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetPatients(c *gin.Context) {
	if c.GetString("role") != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func GetPatient(c *gin.Context) {
	if c.GetString("role") != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	id := c.Param("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func UpdatePatient(c *gin.Context) {
	if c.GetString("role") != "receptionist" {
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

func DeletePatient(c *gin.Context) {
	if c.GetString("role") != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	id := c.Param("id")
	if err := config.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "patient deleted"})
}
