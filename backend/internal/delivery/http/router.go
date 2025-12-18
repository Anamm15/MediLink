package http

import (
	"MediLink/internal/domain/delivery/http/handler"
	"MediLink/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, userHandler handler.UserHandler) {
	user := server.Group("/api/v1/users")
	{
		user.GET("/", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), userHandler.GetAll)
		user.GET("/me", middlewares.Authenticate(), userHandler.GetProfile)

		user.POST("/register", userHandler.Register)
		user.POST("/login", userHandler.Login)
		user.POST("/send-otp", middlewares.Authenticate(), userHandler.SendOTP)
		user.POST("/verify-otp", middlewares.Authenticate(), userHandler.VerifyOTP)
		user.POST("/on-board-patient", middlewares.Authenticate(), userHandler.OnBoardPatient)

		user.PUT("/", middlewares.Authenticate(), userHandler.UpdateProfile)
		user.PATCH("/password", middlewares.Authenticate(), userHandler.ChangePassword)

		user.DELETE("/", middlewares.Authenticate(), userHandler.Delete)
	}
}

func PatientRoute(server *gin.Engine, patientHandler handler.PatientHandler) {
	patient := server.Group("/api/v1/patients")
	{
		patient.PUT("/", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), patientHandler.Update)
	}
}

func ClinicRoute(server *gin.Engine, clinicHandler handler.ClinicHandler) {
	clinic := server.Group("/api/v1/clinics")
	{
		clinic.GET("/", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.GetAll)
		clinic.GET("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.GetByID)
		clinic.GET("/find", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Find)
		clinic.POST("/", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Create)
		clinic.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Update)
		clinic.PATCH("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Update)
		clinic.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Delete)
	}
}
