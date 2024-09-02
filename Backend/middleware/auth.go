package middleware

import (
	"net/http"

	"github.com/adityjoshi/Swaasthya/Backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth token is missing here"})
			c.Abort()
			return
		}
		claims, err := utils.DecodeJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		User_type, ok := claims["user"].(map[string]interface{})["type"].(string)
		if !ok || User_type != "Patient" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized for warden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
