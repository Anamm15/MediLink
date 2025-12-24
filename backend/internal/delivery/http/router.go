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

func DoctorRoute(server *gin.Engine, doctorHandler handler.DoctorHandler) {
	doctor := server.Group("/api/v1/doctors")
	{
		doctor.GET("", doctorHandler.Find)
		doctor.GET("/:id", doctorHandler.GetProfile)
		doctor.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin", "doctor"), doctorHandler.Update)
		doctor.POST("/schedule", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.AddSchedule)
		doctor.PUT("/schedule/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.UpdateSchedule)
	}
}

func MedicineRoute(server *gin.Engine, medicineHandler handler.MedicineHandler) {
	medicine := server.Group("/api/v1/medicines")
	{
		medicine.GET("/", medicineHandler.GetAll)
		medicine.GET("/:id", medicineHandler.GetByID)
		medicine.GET("/search", medicineHandler.Search)
		medicine.POST("/", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Create)
		medicine.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Update)
		medicine.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Delete)
	}
}

func PrescriptionRoute(server *gin.Engine, prescriptionHandler handler.PrescriptionHandler) {
	prescriptions := server.Group("/api/v1/prescriptions")
	{
		prescriptions.GET("/patient", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), prescriptionHandler.GetByPatient)
		prescriptions.GET("/doctor", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.GetByDoctor)
		prescriptions.GET("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("patient", "doctor"), prescriptionHandler.GetDetailByID)
		prescriptions.POST("/", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Create)
		prescriptions.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Update)
		prescriptions.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Delete)
		prescriptions.POST("/:id/add-medicine", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.AddMedicine)
		prescriptions.PATCH("/:id/update-medicine/:prescription_medicine_id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.UpdateMedicine)
		prescriptions.DELETE("/:id/remove-medicine/:prescription_medicine_id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.RemoveMedicine)
	}
}
