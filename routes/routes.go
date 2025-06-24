package routes

import (
	"github.com/gin-gonic/gin"
	"makerble-crudapi/controllers"
)


func SetupRoutes(r *gin.Engine) {
	// Login Route
	r.POST("/login", controllers.LoginHandler)

	// Protected Routes
	auth := r.Group("/api")
	auth.Use(controllers.AuthMiddleware())
	{
		// Receptionist Routes
		auth.POST("/patients", controllers.CreatePatient)
		auth.GET("/patients", controllers.GetPatients)
		auth.GET("/patients/:id", controllers.GetPatient)
		auth.PUT("/patients/:id", controllers.UpdatePatient)
		auth.DELETE("/patients/:id", controllers.DeletePatient)

		// Doctor Routes
		auth.GET("/doctor/patients", controllers.GetAllPatientsForDoctor)
		auth.PUT("/doctor/patients/:id", controllers.UpdatePatientForDoctor)
	}
}