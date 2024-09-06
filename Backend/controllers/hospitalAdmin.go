package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adityjoshi/Swaasthya/Backend/database"
	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	var existingUser database.HospitalAdmin
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
	redisClient := database.GetRedisClient()
	err = redisClient.Set(context.Background(), "otp_verified:"+strconv.Itoa(int(user.AdminID)), "verified", 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error setting OTP verification status"})
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
	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	adminIDUint, ok := adminID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid admin ID"})
		return
	}
	hospital.AdminID = adminIDUint

	var latestHospital database.Hospitals
	if err := database.DB.Order("hospital_id DESC").First(&latestHospital).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve latest hospital"})
			return
		}
	}

	// Generate the hospital username based on HospitalID, HospitalName, and AdminID

	if err := database.DB.Create(&hospital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hospital"})
		return
	}
	hospital.Username = fmt.Sprintf("DEL%d", hospital.HospitalId)
	if err := database.DB.Model(&hospital).Update("username", hospital.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hospital username"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Hospital created successfully", "hospital_id": hospital.HospitalId, "admin": adminIDUint, "username": hospital.Username})
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

func RegisterStaff(c *gin.Context) {
	var staff database.HospitalStaff
	if err := c.BindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get admin ID from JWT
	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	adminIDUint, ok := adminID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid admin ID"})
		return
	}

	// Check if the admin is authorized to register staff for this hospital
	hospitalID, err := verifyAdminHospital(adminIDUint)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin not authorized to register staff"})
		return
	}

	// Set the HospitalID for the new staff
	staff.HospitalID = hospitalID

	// Retrieve hospital details to get the hospital name
	var hospital database.Hospitals
	if err := database.DB.Where("hospital_id = ?", hospitalID).First(&hospital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hospital details"})
		return
	}

	// Set the hospital name in the staff record
	staff.HospitalName = hospital.HospitalName

	// Generate the staff username in the format "hospitalID+staffname"
	staff.Username = fmt.Sprintf("%d%s", hospital.HospitalId, strings.ReplaceAll(strings.ToLower(staff.FullName), " ", ""))

	// Save the new staff entry
	if err := database.DB.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Staff created successfully",
		"staff_id":      staff.StaffID,
		"username":      staff.Username,
		"hospital_id":   staff.HospitalID,
		"hospital_name": staff.HospitalName,
	})
}

func verifyAdminHospital(adminID uint) (uint, error) {
	var admin database.HospitalAdmin
	if err := database.DB.Where("admin_id = ?", adminID).First(&admin).Error; err != nil {
		return 0, err
	}

	var hospital database.Hospitals
	if err := database.DB.Where("admin_id = ?", adminID).First(&hospital).Error; err != nil {
		return 0, err
	}

	return hospital.HospitalId, nil
}
