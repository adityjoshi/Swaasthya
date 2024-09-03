package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/gin-gonic/gin"
)

// CreateAppointment handles POST requests to create a new appointment
func CreateAppointment(c *gin.Context) {
	var appointmentData struct {
		PatientID       uint      `json:"patient_id"`
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
		PatientID:       appointmentData.PatientID,
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

	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "appointment_id": appointment.AppointmentID})
}
