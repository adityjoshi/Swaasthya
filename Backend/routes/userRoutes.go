package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register")
	incomingRoutes.POST("/login")
}
