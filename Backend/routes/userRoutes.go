package routes

import (
	"github.com/adityjoshi/Swaasthya/Backend/controllers"
	"github.com/adityjoshi/Swaasthya/Backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.Register)
	incomingRoutes.POST("/login", controllers.Login)
	incomingRoutes.POST("/verify-otp", controllers.VerifyOTP)
	incomingRoutes.POST("/bookAppointment", middleware.AuthRequired("Patient"), controllers.CreateAppointment)
}
