package controllers

import (
	"log"

	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/go-redis/redis/v8"
)

// GenerateAndSendOTP generates an OTP, stores it in Redis, and sends it via email.
func GenerateAndSendOTP(email string) (string, error) {
	// Generate OTP
	otp, err := utils.GenerateOtp()
	if err != nil {
		return "", err
	}

	// Store OTP in Redis with an expiration time
	err = utils.StoreOtp(email+"_otp", otp)
	if err != nil {
		return "", err
	}

	// Send OTP to user via email asynchronously
	go func() {
		err := utils.OtpRegistration(email, otp)
		if err != nil {
			log.Printf("Failed to send OTP email to %s: %v", email, err)
		} else {
			log.Printf("Successfully sent OTP to %s", email)
		}
	}()

	return otp, nil
}

// VerifyOtp verifies the provided OTP against the stored OTP.
func VerifyOtp(email, otp string) (bool, error) {
	storedOtp, err := utils.GetOtp(email + "_otp")
	if err == redis.Nil {
		log.Printf("OTP not found for email: %s", email)
		return false, nil
	} else if err != nil {
		return false, err
	}

	if otp != storedOtp {
		log.Printf("OTP mismatch for email: %s", email)
		return false, nil
	}

	// Delete OTP after successful verification
	err = utils.DeleteOTP(email + "_otp")
	if err != nil {
		log.Printf("Failed to delete OTP for email: %s", email)
		return false, err
	}

	return true, nil
}
