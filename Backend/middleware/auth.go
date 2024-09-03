package middleware

import (
	"net/http"
	"strings"

	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthRequired(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		// Remove 'Bearer ' prefix if present
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[len("Bearer "):]
		}

		claims, err := utils.DecodeJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check user type
		tokenUserType, ok := claims["user_type"].(string)
		if !ok || tokenUserType != userType {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Set user ID in context
		userID, _ := claims["user_id"].(float64)
		c.Set("user_id", uint(userID))

		c.Next()
	}
}
