// package utils

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/adityjoshi/Swaasthya/Backend/database"
// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtSecret = []byte(os.Getenv("JWTSECRET"))

// func GenerateJwt(user_id int, userType database.Users) (string, error) {
// 	// Create a new JWT token with user payload
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user"] = map[string]interface{}{
// 		"userid":    user_id,
// 		"User_type": userType,
// 	}
// 	claims["exp"] = time.Now().Add(time.Hour).Unix()

// 	// sign the token with the secret key and return
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil

// }

// func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Extract and return the claims if the token is valid
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	}
// 	return nil, fmt.Errorf("invalid token")

// }

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWTSECRET"))

// GenerateJwt generates a JWT token with the user ID and User_type as claims.
func GenerateJwt(user_id int, userType string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Add claims to the token
	claims["user"] = map[string]interface{}{
		"userid":    user_id,
		"User_type": userType,
	}
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Set expiration time

	// Sign the token with the secret key and return
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeJwt decodes a JWT token and returns the claims.
func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract and return the claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
