package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StaffLogin(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var staff database.HospitalStaff
	if err := database.DB.Where("email = ?", loginRequest.Email).First(&staff).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate and send OTP
	otp, err := GenerateAndSendOTP(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate or send OTP" + otp})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJwt(staff.StaffID, "Staff")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with message to enter OTP
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email. Please verify the OTP.", "token": token})
}
func VerifyStaffOTP(c *gin.Context) {
	var otpRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.BindJSON(&otpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify the OTP
	isValid, err := VerifyOtp(otpRequest.Email, otpRequest.OTP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error verifying OTP"})
		return
	}
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	// Retrieve user information after OTP verification
	var staff database.HospitalStaff
	if err := database.DB.Where("email = ?", otpRequest.Email).First(&staff).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Staff not found"})
		return
	}

	// Set OTP verification status in Redis
	redisClient := database.GetRedisClient()
	err = redisClient.Set(context.Background(), "otp_verified:"+strconv.Itoa(int(staff.StaffID)), "verified", 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error setting OTP verification status"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	claims, err := utils.DecodeJwt(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
		return
	}

	// Extract user_type from JWT claims
	userType, ok := claims["user_type"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loggedin": "success", "user": userType, "staff": database.Staff})
}

func RegisterPatient(c *gin.Context) {
	var patient database.Patients
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get staff ID from JWT or context
	staffID, exists := c.Get("staff_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	staffIDUint, ok := staffID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Retrieve hospital ID associated with the staff
	var staff database.HospitalStaff
	if err := database.DB.Where("staff_id = ?", staffIDUint).First(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve staff details"})
		return
	}

	// Set the HospitalID based on the staff's hospital
	patient.HospitalID = staff.HospitalID

	// Create a new patient entry in the database
	if err := database.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Patient registered successfully",
		"patient": patient,
	})
}
