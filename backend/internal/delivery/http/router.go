package http

import (
	"MediLink/internal/domain/delivery/http/handler"
	"MediLink/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoute(server *gin.Engine, authHandler handler.AuthHandler) {
	auth := server.Group("/api/v1/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
		auth.POST("/refresh-token", authHandler.RefreshToken)
		auth.POST("/logout", middlewares.Authenticate(), authHandler.Logout)
		auth.POST("/change-password", middlewares.Authenticate(), authHandler.ChangePassword)
		auth.POST("/request-reset-password", authHandler.RequestResetPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}
}

func UserRoute(server *gin.Engine, userHandler handler.UserHandler) {
	user := server.Group("/api/v1/users")
	{
		user.GET("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), userHandler.GetAll)
		user.GET("/me", middlewares.Authenticate(), userHandler.Me)
		user.GET("/profile", middlewares.Authenticate(), userHandler.GetProfile)

		user.POST("/send-email-verification", middlewares.Authenticate(), userHandler.SendVerificationUser)
		user.POST("/verify-email", middlewares.Authenticate(), userHandler.VerifyUser)
		user.POST("/onboard-patient", middlewares.Authenticate(), userHandler.OnBoardPatient)
		user.PUT("", middlewares.Authenticate(), userHandler.UpdateProfile)
		user.DELETE("", middlewares.Authenticate(), userHandler.Delete)
	}
}

func PatientRoute(server *gin.Engine, patientHandler handler.PatientHandler) {
	patient := server.Group("/api/v1/patients")
	{
		patient.GET("/me", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), patientHandler.Me)
		patient.PUT("", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), patientHandler.Update)
	}
}

func ClinicRoute(server *gin.Engine, clinicHandler handler.ClinicHandler, clinicInventoryHandler handler.ClinicInventoryHandler) {
	clinics := server.Group("/api/v1/clinics")
	{
		clinics.GET("", clinicHandler.GetAll)
		clinics.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Create)

		clinic := clinics.Group("/:clinic_id")
		{
			clinic.GET("", clinicHandler.GetByID)
			clinic.PUT("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Update)
			clinic.DELETE("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.Delete)

			clinic.POST("/doctors", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.AssignDoctor)
			clinic.DELETE("/doctors/:doctor_id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicHandler.RemoveDoctor)

			inventories := clinic.Group("/inventories")
			{
				inventories.GET("", clinicInventoryHandler.GetByClinic)
				inventories.GET("/:id", clinicInventoryHandler.GetByID)
				inventories.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicInventoryHandler.Create)
				inventories.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicInventoryHandler.Update)
				inventories.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), clinicInventoryHandler.Delete)
			}
		}
	}
}

func DoctorRoute(server *gin.Engine, doctorHandler handler.DoctorHandler) {
	doctor := server.Group("/api/v1/doctors")
	{
		doctor.GET("/search", doctorHandler.Find)
		doctor.GET("/me", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.Me)
		doctor.GET("/:id", doctorHandler.GetProfile)
		doctor.PUT("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin", "doctor"), doctorHandler.Update)

		schedule := doctor.Group("/schedules")
		{
			schedule.GET("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin", "doctor"), doctorHandler.GetSchedules)
			schedule.GET("/:id", doctorHandler.GetScheduleByID)
			schedule.GET("/availability", doctorHandler.GetAvailableSchedules)
			schedule.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.AddSchedule)
			schedule.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.UpdateSchedule)
			schedule.PATCH("/:id/status", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.UpdateScheduleStatus)
			schedule.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), doctorHandler.DeleteSchedule)
		}
	}
}

func MedicineRoute(server *gin.Engine, medicineHandler handler.MedicineHandler) {
	medicine := server.Group("/api/v1/medicines")
	{
		medicine.GET("", medicineHandler.GetAll)
		medicine.GET("/:id", medicineHandler.GetByID)
		medicine.GET("/search", medicineHandler.Search)
		medicine.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Create)
		medicine.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Update)
		medicine.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), medicineHandler.Delete)
	}
}

func AppointmentRoute(server *gin.Engine, appointmentHandler handler.AppointmentHandler) {
	appointment := server.Group("/api/v1/appointments")
	{
		appointment.GET("", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), appointmentHandler.GetAll)
		appointment.GET("/:id", middlewares.Authenticate(), appointmentHandler.GetDetailByID)
		appointment.GET("/doctor", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), appointmentHandler.GetByDoctor)
		appointment.GET("/patient", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), appointmentHandler.GetByPatient)
		appointment.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("patient", "doctor", "admin"), appointmentHandler.Create)
		appointment.PATCH("/:id/cancel", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), appointmentHandler.CancelBooking)
		appointment.PATCH("/:id/complete", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), appointmentHandler.CompleteConsultation)
		appointment.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), appointmentHandler.Delete)
	}
}

func MedicalRecordRoute(server *gin.Engine, medicalRecordHandler handler.MedicalRecordHandler) {
	medicalRecord := server.Group("/api/v1/medical-records")
	{
		medicalRecord.GET("/patient/:patient_id", middlewares.Authenticate(), middlewares.AuthorizeRole("patient", "admin"), medicalRecordHandler.GetByPatient)
		medicalRecord.GET("/doctor/:doctor_id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor", "admin"), medicalRecordHandler.GetByDoctor)
		medicalRecord.GET("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("patient", "doctor", "admin"), medicalRecordHandler.GetDetailByID)
		medicalRecord.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), medicalRecordHandler.Create)
		medicalRecord.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), medicalRecordHandler.Update)
		medicalRecord.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor", "patient"), medicalRecordHandler.Delete)
	}
}

func PrescriptionRoute(server *gin.Engine, prescriptionHandler handler.PrescriptionHandler) {
	prescriptions := server.Group("/api/v1/prescriptions")
	{
		prescriptions.GET("/patient", middlewares.Authenticate(), middlewares.AuthorizeRole("patient"), prescriptionHandler.GetByPatient)
		prescriptions.GET("/doctor", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.GetByDoctor)
		prescriptions.GET("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("patient", "doctor"), prescriptionHandler.GetDetailByID)
		prescriptions.POST("", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Create)
		prescriptions.PUT("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Update)
		prescriptions.DELETE("/:id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.Delete)
		prescriptions.POST("/:id/add-medicine", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.AddMedicine)
		prescriptions.PATCH("/:id/update-medicine/:prescription_medicine_id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.UpdateMedicine)
		prescriptions.DELETE("/:id/remove-medicine/:prescription_medicine_id", middlewares.Authenticate(), middlewares.AuthorizeRole("doctor"), prescriptionHandler.RemoveMedicine)
	}
}

func PaymentRoute(server *gin.Engine, paymentHandler handler.PaymentHandler) {
	payments := server.Group("/api/v1/payments")
	{
		payments.POST("/callback", paymentHandler.ReceiveNotification)
	}
}
