package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adityjoshi/Swaasthya/Backend/database"

	"github.com/gin-gonic/gin"
)

// func RegisterDoctor(c *gin.Context) {
// 	var doctorData struct {
// 		FullName      string              `json:"full_name"`
// 		Description   string              `json:"description"`
// 		ContactNumber string              `json:"contact_number"`
// 		Email         string              `json:"email"`
// 		AdminID       uint                `json:"admin_id"`
// 		Department    database.Department `json:"department"`
// 	}

// 	if err := c.BindJSON(&doctorData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	// Ensure AdminID is included in the JSON payload
// 	if doctorData.AdminID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin ID is required"})
// 		return
// 	}

// 	// Find the hospital associated with the admin
// 	var hospital database.Hospitals
// 	if err := database.DB.Where("admin_id = ?", doctorData.AdminID).First(&hospital).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Hospital not found for the admin"})
// 		return
// 	}

// 	// Set HospitalID and HospitalName in doctor data
// 	doctor := database.Doctors{
// 		FullName:      doctorData.FullName,
// 		Description:   doctorData.Description,
// 		ContactNumber: doctorData.ContactNumber,
// 		Email:         doctorData.Email,
// 		HospitalID:    hospital.HospitalId,   // Correctly set HospitalID from fetched hospital
// 		Hospital:      hospital.HospitalName, // Set HospitalName
// 		Department:    doctorData.Department,
// 	}

// 	// Generate username
// 	doctor.Username = generateDoctorUsername(doctor.HospitalID, hospital.HospitalName, doctor.FullName)

// 	// Save doctor data
// 	if err := database.DB.Create(&doctor).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register doctor"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Doctor registered successfully", "hospital_name": hospital.HospitalName})
// }

func RegisterDoctor(c *gin.Context) {
	var doctorData struct {
		FullName      string              `json:"full_name"`
		Description   string              `json:"description"`
		ContactNumber string              `json:"contact_number"`
		Email         string              `json:"email"`
		Department    database.Department `json:"department"`
	}

	if err := c.BindJSON(&doctorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Extract AdminID from JWT claims
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	adminIDUint, ok := adminID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid admin ID"})
		return
	}

	// Find the hospital associated with the admin
	var hospital database.Hospitals
	if err := database.DB.Where("admin_id = ?", adminIDUint).First(&hospital).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hospital not found for the admin"})
		return
	}

	// Set HospitalID and HospitalName in doctor data
	doctor := database.Doctors{
		FullName:      doctorData.FullName,
		Description:   doctorData.Description,
		ContactNumber: doctorData.ContactNumber,
		Email:         doctorData.Email,
		HospitalID:    hospital.HospitalId,   // Correctly set HospitalID from fetched hospital
		Hospital:      hospital.HospitalName, // Set HospitalName
		Department:    doctorData.Department,
	}

	// Generate username
	doctor.Username = generateDoctorUsername(doctor.HospitalID, hospital.HospitalName, doctor.FullName)

	// Save doctor data
	if err := database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor registered successfully", "hospital_name": hospital.HospitalName})
}

// Helper function to generate doctor username
func generateDoctorUsername(hospitalID uint, hospitalName, doctorFullName string) string {
	// Remove spaces from hospital name and doctor full name
	hospitalName = strings.ReplaceAll(hospitalName, " ", "")
	doctorFullName = strings.ReplaceAll(doctorFullName, " ", "")
	// Construct username
	return fmt.Sprintf("%d%s%s", hospitalID, hospitalName, doctorFullName)
}

func GetDoctor(c *gin.Context) {
	doctorID := c.Param("doctor_id")

	var doctor database.Doctors
	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	var hospital database.Hospitals
	if err := database.DB.First(&hospital, doctor.HospitalID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hospital not found for the doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor_id":      doctor.DoctorID,
		"full_name":      doctor.FullName,
		"description":    doctor.Description,
		"contact_number": doctor.ContactNumber,
		"email":          doctor.Email,
		"hospital_id":    doctor.HospitalID,
		"hospital_name":  hospital.HospitalName,
		"department":     doctor.Department,
		"username":       doctor.Username,
	})
}
