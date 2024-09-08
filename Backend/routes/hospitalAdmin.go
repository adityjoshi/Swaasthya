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
	incomingRoutes.POST("/stafflogin", controllers.StaffLogin)
	incomingRoutes.POST("/staffotp", controllers.VerifyStaffOTP)
	incomingRoutes.POST("/pati", middleware.AuthRequired("Staff", "Reception"), controllers.RegisterPatient)
	incomingRoutes.POST("/admit", middleware.AuthRequired("Staff", "Reception"), controllers.AdmitPatientForHospitalization)
	incomingRoutes.POST("/compounder", controllers.CompounderLogin)
	incomingRoutes.POST("/markCompounder", middleware.AuthRequired("Staff", "Compounder"), controllers.MarkPatientAsHospitalized)
	incomingRoutes.GET("/get", middleware.AuthRequired("Staff", "Compounder"), controllers.GetRoomAssignments)
	// incomingRoutes.POST("/registerhospital", controllers.RegisterHospital)
	// incomingRoutes.GET("/gethospital/:id", controllers.GetHospital)
	// incomingRoutes.POST("/doctor", controllers.RegisterDoctor)
	// incomingRoutes.GET("/getdoctor/:id", controllers.GetDoctor)
	// incomingRoutes.POST("/bookAppointment", controllers.CreateAppointment)

	adminRoutes := incomingRoutes.Group("/admin")
	adminRoutes.Use(middleware.AuthRequired("Admin", ""))
	{
		adminRoutes.POST("/registerhospital", middleware.OtpAuthRequireed, controllers.RegisterHospital)
		adminRoutes.GET("/gethospital/:id", middleware.OtpAuthRequireed, controllers.GetHospital)
		adminRoutes.POST("/doctor", middleware.OtpAuthRequireed, controllers.RegisterDoctor)
		adminRoutes.GET("/getdoctor/:id", middleware.OtpAuthRequireed, controllers.GetDoctor)
		adminRoutes.POST("/bookAppointment", middleware.OtpAuthRequireed, controllers.CreateAppointment)
		adminRoutes.POST("/registerStaff", middleware.OtpAuthRequireed, controllers.RegisterStaff)
		adminRoutes.POST("/registerBeds", middleware.OtpAuthRequireed, controllers.AddBedType)
		adminRoutes.POST("/updateBeds", middleware.OtpAuthRequireed, controllers.UpdateTotalBeds)
		adminRoutes.GET("/getBeds", middleware.OtpAuthRequireed, controllers.GetTotalBeds)
	}
}
