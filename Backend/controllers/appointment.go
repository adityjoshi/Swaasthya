package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
)

// // CreateAppointment handles POST requests to create a new appointment
// func CreateAppointment(c *gin.Context) {
// 	var appointmentData struct {
// 		PatientID       uint      `json:"patient_id"`
// 		DoctorID        uint      `json:"doctor_id"`
// 		AppointmentDate time.Time `json:"appointment_date"`
// 		AppointmentTime time.Time `json:"appointment_time"`
// 		Description     string    `json:"description"`
// 	}

// 	if err := c.BindJSON(&appointmentData); err != nil {
// 		fmt.Println("Error binding JSON:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
// 		return
// 	}

// 	appointment := database.Appointment{
// 		PatientID:       appointmentData.PatientID,
// 		DoctorID:        appointmentData.DoctorID,
// 		AppointmentDate: appointmentData.AppointmentDate,
// 		AppointmentTime: appointmentData.AppointmentTime,
// 		Description:     appointmentData.Description,
// 	}

// 	if err := database.DB.Create(&appointment).Error; err != nil {
// 		fmt.Println("Error creating appointment:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
// 		return
// 	}
// 	var doctor database.Doctors
// 	if err := database.DB.Where("doctor_id = ?", appointmentData.DoctorID).First(&doctor).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor information"})
// 		return
// 	}

// 	// Fetch patient details
// 	var patient database.Users
// 	if err := database.DB.Where("patient_id = ?", appointmentData.PatientID).First(&patient).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient information"})
// 		return
// 	}

// 	// Send appointment email to the patient
// 	bookingTime := time.Now().Format("2006-01-02 15:04:05") // Current time as booking time
// 	err := utils.SendAppointmentEmail(patient.Email, doctor.FullName, appointmentData.AppointmentDate.Format("2006-01-02"), appointmentData.AppointmentTime.Format("15:04"), bookingTime)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send appointment email"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "appointment_id": appointment.AppointmentID})
// }

func CreateAppointment(c *gin.Context) {
	var appointmentData struct {
		UserID          uint      `json:"user_id"`
		DoctorID        uint      `json:"doctor_id"`
		AppointmentDate time.Time `json:"appointment_date"`
		AppointmentTime time.Time `json:"appointment_time"`
		Description     string    `json:"description"`
	}

	if err := c.BindJSON(&appointmentData); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	appointment := database.Appointment{
		PatientID:       appointmentData.UserID, // Assuming PatientID is the foreign key for Users table
		DoctorID:        appointmentData.DoctorID,
		AppointmentDate: appointmentData.AppointmentDate,
		AppointmentTime: appointmentData.AppointmentTime,
		Description:     appointmentData.Description,
	}

	if err := database.DB.Create(&appointment).Error; err != nil {
		fmt.Println("Error creating appointment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	var doctor database.Doctors
	if err := database.DB.Where("doctor_id = ?", appointmentData.DoctorID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor information"})
		return
	}

	// Fetch user details
	var user database.Users
	if err := database.DB.Where("user_id = ?", appointmentData.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
		return
	}

	// Send appointment email to the user
	bookingTime := time.Now().Format("2006-01-02 15:04:05") // Current time as booking time
	err := utils.SendAppointmentEmail(user.Email, doctor.FullName, appointmentData.AppointmentDate.Format("2006-01-02"), appointmentData.AppointmentTime.Format("15:04"), bookingTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send appointment email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "appointment_id": appointment.AppointmentID})
}
