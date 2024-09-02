package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/adityjoshi/Swaasthya/Backend/database"
)

func GenerateOtp() (string, error) {
	otp, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", otp.Int64()), nil
}

func StoreOtp(key, otp string) error {
	client := database.GetRedisClient()
	// otp will expire after 5 min
	return client.Set(database.Ctx, key, otp, 5*time.Minute).Err()
}

// Retrieve OTP from Redis
func GetOtp(key string) (string, error) {
	client := database.GetRedisClient()

	otp, err := client.Get(database.Ctx, key).Result()
	if err != nil {
		return "", err
	}
	return otp, nil
}

// Delete OTP from Redis
func DeleteOTP(key string) error {
	client := database.GetRedisClient()
	return client.Del(database.Ctx, key).Err()
}
