package routes

import (
	"time"

	"github.com/adityjoshi/Swaasthya/Backend/controllers"
	"github.com/adityjoshi/Swaasthya/Backend/middleware"
	"github.com/gin-gonic/gin"
)

func HospitalAdmin(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/hospitaladmin", controllers.RegisterHospitalAdmin)
	incomingRoutes.POST("/adminLogin", middleware.RateLimiterMiddleware(2, time.Minute), controllers.AdminLogin)
	incomingRoutes.POST("/adminOtp", controllers.VerifyAdminOTP)
	// incomingRoutes.POST("/registerhospital", controllers.RegisterHospital)
	// incomingRoutes.GET("/gethospital/:id", controllers.GetHospital)
	// incomingRoutes.POST("/doctor", controllers.RegisterDoctor)
	// incomingRoutes.GET("/getdoctor/:id", controllers.GetDoctor)
	// incomingRoutes.POST("/bookAppointment", controllers.CreateAppointment)

	adminRoutes := incomingRoutes.Group("/admin")
	adminRoutes.Use(middleware.AuthRequired("Admin"))
	{
		adminRoutes.POST("/registerhospital", middleware.OtpAuthRequireed, controllers.RegisterHospital)
		adminRoutes.GET("/gethospital/:id", middleware.OtpAuthRequireed, controllers.GetHospital)
		adminRoutes.POST("/doctor", middleware.OtpAuthRequireed, controllers.RegisterDoctor)
		adminRoutes.GET("/getdoctor/:id", middleware.OtpAuthRequireed, controllers.GetDoctor)
		adminRoutes.POST("/bookAppointment", middleware.OtpAuthRequireed, controllers.CreateAppointment)
		adminRoutes.POST("/registerStaff", middleware.OtpAuthRequireed, controllers.RegisterStaff)
	}
}
