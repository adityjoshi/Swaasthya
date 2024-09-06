package controllers

import (
	"fmt"
	"net/http"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/gin-gonic/gin"
)

// AddPatientDetails handles adding patient details
func AddPatientDetails(c *gin.Context) {
	// Extract patient_id from URL parameters
	patientID := c.Param("id")

	// Bind JSON body to patientInfo
	var patientInfo database.PatientInfo
	if err := c.BindJSON(&patientInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert patientID to uint if needed (assuming PatientID in database is of type uint)
	var id uint
	if _, err := fmt.Sscan(patientID, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	// Set the PatientID from the URL parameter
	patientInfo.PatientID = id

	// Fetch user details based on PatientID
	var user database.Users
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate PatientUsername
	patientUsername := fmt.Sprintf(user.Full_Name + user.ContactNumber)
	patientInfo.Username = patientUsername

	// Create patientInfo record
	if err := database.DB.Create(&patientInfo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save patient information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient information added successfully", "patient_username": patientUsername})
}

func GetPatientDetails(c *gin.Context) {
	// Extract patient_id from URL parameters
	patientID := c.Param("id")

	// Fetch patient details based on PatientID
	var patientInfo database.PatientInfo
	if err := database.DB.Where("patient_id = ?", patientID).First(&patientInfo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient information not found"})
		return
	}

	c.JSON(http.StatusOK, patientInfo)
}
