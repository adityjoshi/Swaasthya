package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHospitalAdmin(c *gin.Context) {
	var admin database.HospitalAdmin

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if admin.Usertype == "" {
		admin.Usertype = "Admin" // Default to "Admin" if not provided
	}
	var existingUser database.Users
	if err := database.DB.Where("email = ?", admin.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	admin.Password = string(hashedPassword)

	if err := database.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hospital admin registered successfully", "admin_id": admin.AdminID})
}

func AdminLogin(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var admin database.HospitalAdmin
	if err := database.DB.Where("email = ?", loginRequest.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}
	otp, err := GenerateAndSendOTP(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate or send OTP" + otp})
		return
	}

	// token, err := utils.GenerateJwt(int(admin.AdminID), "Admin")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	// 	return
	// }

	token, err := utils.GenerateJwt(admin.AdminID, "Admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with message to enter OTP
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email. Please verify the OTP.", "token": token})
}
func VerifyAdminOTP(c *gin.Context) {
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
	var user database.HospitalAdmin
	if err := database.DB.Where("email = ?", otpRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loggedin": "success"})
}

func RegisterHospital(c *gin.Context) {
	var hospital database.Hospitals
	if err := c.BindJSON(&hospital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adminIDUint := hospital.AdminID

	hospital.HospitalName = strings.TrimSpace(hospital.HospitalName)

	// Generate the hospital username based on HospitalID, HospitalName, and AdminID

	if err := database.DB.Create(&hospital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hospital"})
		return
	}
	hospital.Username = fmt.Sprintf("HOSP-%d-%s-ADMIN-%d", hospital.HospitalId, strings.ReplaceAll(hospital.HospitalName, " ", ""), adminIDUint)
	c.JSON(http.StatusCreated, gin.H{"message": "Hospital created successfully", "hospital_id": hospital.HospitalId, "admin": adminIDUint})
}

func GetHospital(c *gin.Context) {
	hospitalID := c.Param("hospital_id")

	var hospital database.Hospitals
	if err := database.DB.First(&hospital, hospitalID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hospital not found"})
		return
	}

	c.JSON(http.StatusOK, hospital)
}
