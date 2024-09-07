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

	token, err := utils.GenerateJwt(admin.AdminID, "Admin", "")
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

	// Verify admin's hospital authorization
	hospitalID, err := verifyAdminHospital(adminIDUint)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin not authorized to register staff"})
		return
	}

	staff.HospitalID = hospitalID

	// Retrieve hospital details
	var hospital database.Hospitals
	if err := database.DB.Where("hospital_id = ?", hospitalID).First(&hospital).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hospital details"})
		return
	}
	staff.HospitalName = hospital.HospitalName

	// Generate a password based on staff's full name and hospital username
	password := generatePassword(staff.FullName, hospital.Username)
	staff.Password = password

	// Hash the staff password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	staff.Password = string(hashedPassword)

	// Generate a unique username for the staff
	staff.Username = fmt.Sprintf("%d%s", hospital.HospitalId, strings.ReplaceAll(strings.ToLower(staff.FullName), " ", ""))

	// Save the new staff entry
	if err := database.DB.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Staff created successfully",
		"staff_id":    staff.StaffID,
		"username":    staff.Username,
		"hospital_id": staff.HospitalID,
		"password":    staff.Password, // Optionally return the generated password
	})
}

func generatePassword(fullName string, hospitalUsername string) string {
	cleanedName := strings.ReplaceAll(strings.ToLower(fullName), " ", "")
	return fmt.Sprintf("%s%s", cleanedName, hospitalUsername)
}

func AddBedType(c *gin.Context) {
	var bedsCount database.BedsCount

	// Parse the JSON request body into the bedsCount struct
	if err := c.BindJSON(&bedsCount); err != nil {
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

	// Verify the admin's hospital
	hospitalID, err := verifyAdminHospital(adminIDUint)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin not authorized to add beds for this hospital"})
		return
	}

	// Set the hospital ID to the bedsCount object
	bedsCount.HospitalID = hospitalID

	// Check if the bed type already exists for the hospital
	var existingBedType database.BedsCount
	if err := database.DB.Where("hospital_id = ? AND type_name = ?", hospitalID, bedsCount.TypeName).First(&existingBedType).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Bed type already exists for this hospital"})
		return
	}

	// Save the new bed type and total beds
	if err := database.DB.Create(&bedsCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add bed type"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Bed type added successfully",
		"bed_type_id": bedsCount.ID,
		"type_name":   bedsCount.TypeName,
		"total_beds":  bedsCount.TotalBeds,
		"hospital_id": bedsCount.HospitalID,
	})
}

func UpdateTotalBeds(c *gin.Context) {
	var bedData struct {
		TypeName  string `json:"type_name"`
		TotalBeds int    `json:"total_beds"` // Number of beds to add or remove
		Action    string `json:"action"`     // Action: "add" or "remove"
	}

	// Parse the JSON request body into the bedData struct
	if err := c.BindJSON(&bedData); err != nil {
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

	// Verify the admin's hospital
	hospitalID, err := verifyAdminHospital(adminIDUint)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin not authorized to update beds for this hospital"})
		return
	}

	// Find the bed type for the given hospital
	var bedType database.BedsCount
	if err := database.DB.Where("hospital_id = ? AND type_name = ?", hospitalID, bedData.TypeName).First(&bedType).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bed type not found for this hospital"})
		return
	}

	// Perform the add or remove action
	switch bedData.Action {
	case "add":
		// Add beds
		previousTotalBeds := bedType.TotalBeds
		bedType.TotalBeds += uint(bedData.TotalBeds)

		// Add rooms for the newly added beds
		for i := previousTotalBeds + 1; i <= bedType.TotalBeds; i++ {
			roomNumber := fmt.Sprintf("%s%d", strings.ToLower(bedData.TypeName), i)
			newRoom := database.Room{
				HospitalID: hospitalID,
				BedType:    bedData.TypeName,
				RoomNumber: roomNumber,
				IsOccupied: false,
			}
			if err := database.DB.Create(&newRoom).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create rooms"})
				return
			}
		}

	case "remove":
		// Remove beds (ensure total beds don't go below zero)
		if int(bedType.TotalBeds)-bedData.TotalBeds < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot remove more beds than available"})
			return
		}

		// Remove unoccupied rooms first
		var unoccupiedRooms []database.Room
		if err := database.DB.Where("hospital_id = ? AND bed_type = ? AND is_occupied = ?", hospitalID, bedData.TypeName, false).Order("room_number desc").Limit(bedData.TotalBeds).Find(&unoccupiedRooms).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find unoccupied rooms"})
			return
		}

		// Ensure there are enough unoccupied rooms to remove
		if len(unoccupiedRooms) < bedData.TotalBeds {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough unoccupied rooms to remove"})
			return
		}

		// Delete the unoccupied rooms
		for _, room := range unoccupiedRooms {
			if err := database.DB.Delete(&room).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room"})
				return
			}
		}

		// Update total beds
		bedType.TotalBeds -= uint(bedData.TotalBeds)

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'add' or 'remove'"})
		return
	}

	// Save the updated bed count
	if err := database.DB.Save(&bedType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bed count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Total beds updated successfully",
		"type_name":  bedType.TypeName,
		"total_beds": bedType.TotalBeds,
	})
}

func GetTotalBeds(c *gin.Context) {
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

	// Verify the admin's hospital
	hospitalID, err := verifyAdminHospital(adminIDUint)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin not authorized to view beds for this hospital"})
		return
	}

	// Query the BedsCount for the given hospital
	var beds []database.BedsCount
	if err := database.DB.Where("hospital_id = ?", hospitalID).Find(&beds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bed information"})
		return
	}

	// If no beds are found for the hospital
	if len(beds) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No bed data found for this hospital"})
		return
	}

	// Create a response that lists all the bed types with their total, available, and occupied counts
	var bedDetails []gin.H
	for _, bed := range beds {
		bedDetails = append(bedDetails, gin.H{
			"type_name":  bed.TypeName,
			"total_beds": bed.TotalBeds,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"hospital_id": hospitalID,
		"bed_details": bedDetails,
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
