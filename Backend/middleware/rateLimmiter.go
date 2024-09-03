package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adityjoshi/Swaasthya/Backend/database"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := database.GetRedisClient()
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		// Increment the request count for the IP
		count, err := client.Incr(database.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Set expiration time for the key if it's the first request
		if count == 1 {
			client.Expire(database.Ctx, key, duration)
		}

		// Check if the limit is exceeded
		if int(count) > limit {
			ttl, _ := client.TTL(database.Ctx, key).Result()
			c.Header("Retry-After", strconv.Itoa(int(ttl.Seconds())))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Try again later.", "ip": c.ClientIP()})
			return
		}

		// Proceed with the request
		c.Next()
	}
}
