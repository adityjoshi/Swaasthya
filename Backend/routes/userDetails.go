package routes

import (
	"github.com/adityjoshi/Swaasthya/Backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserInfoRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/updatePatientInfo/:id", controllers.AddPatientDetails)
	// incomingRoutes.GET("/getPatientId/:id", middleware.AuthUser(), controllers.GetPatientDetails)

}
