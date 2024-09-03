package routes

import (
	"github.com/adityjoshi/Swaasthya/Backend/controllers"
	"github.com/adityjoshi/Swaasthya/Backend/middleware"
	"github.com/gin-gonic/gin"
)

func HospitalAdmin(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/hospitaladmin", controllers.RegisterHospitalAdmin)
	incomingRoutes.POST("/adminLogin", controllers.AdminLogin)
	incomingRoutes.POST("/adminOtp", controllers.VerifyAdminOTP)
	// incomingRoutes.POST("/registerhospital", controllers.RegisterHospital)
	// incomingRoutes.GET("/gethospital/:id", controllers.GetHospital)
	// incomingRoutes.POST("/doctor", controllers.RegisterDoctor)
	// incomingRoutes.GET("/getdoctor/:id", controllers.GetDoctor)
	// incomingRoutes.POST("/bookAppointment", controllers.CreateAppointment)

	adminRoutes := incomingRoutes.Group("/admin")
	adminRoutes.Use(middleware.AuthRequired("Admin"))
	{
		adminRoutes.POST("/registerhospital", controllers.RegisterHospital)
		adminRoutes.GET("/gethospital/:id", controllers.GetHospital)
		adminRoutes.POST("/doctor", controllers.RegisterDoctor)
		adminRoutes.GET("/getdoctor/:id", controllers.GetDoctor)
		adminRoutes.POST("/bookAppointment", controllers.CreateAppointment)
	}
}
