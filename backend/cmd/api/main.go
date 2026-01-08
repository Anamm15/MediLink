package main

import (
	"log"
	"os"

	"MediLink/internal/delivery/http"
	"MediLink/internal/delivery/http/handler"
	"MediLink/internal/infrastructure/database"
	"MediLink/internal/infrastructure/repository"
	"MediLink/internal/middlewares"
	"MediLink/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")

	server := gin.Default()
	server.Use(middlewares.SetupCORS())

	db := database.SetUpDatabaseConnection()
	redisClient := database.NewRedisClient()

	userRepository := repository.NewUserRepository(db)
	authRepository := repository.NewRefreshTokenRepository(db)
	patientRepository := repository.NewPatientRepository(db)
	redisRepository := repository.NewRedisRepository(redisClient)
	clinicRepository := repository.NewClinicRepository(db)
	doctorRepository := repository.NewDoctorRepository(db)
	doctorClinicReplacementRepository := repository.NewDoctorClinicPlacementRepository(db)
	doctorScheduleRepository := repository.NewDoctorScheduleRepository(db)
	medicineRepository := repository.NewMedicineRepository(db)
	clinicInventoryRepository := repository.NewClinicInventoryRepository(db)
	appointmentRepository := repository.NewAppointmentRepository(db)
	medicalRecordRepository := repository.NewMedicalRecordRepository(db)
	prescriptionRepository := repository.NewPrescriptionRepository(db)
	prescriptionItemRepository := repository.NewPrescriptionItemRepository(db)
	billingRepository := repository.NewBillingRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)

	notificationUsecase := usecase.NewNotificationUsecase(redisRepository)
	paymentUsecase := usecase.NewPaymentUsecase(db, midtransServerKey, paymentRepository, billingRepository, appointmentRepository)
	userUsecase := usecase.NewUserUsecase(userRepository, patientRepository, redisRepository, notificationUsecase)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository, redisRepository, notificationUsecase)
	patientUsecase := usecase.NewPatientUsecase(patientRepository, redisRepository)
	clinicUsecase := usecase.NewClinicUsecase(clinicRepository, doctorClinicReplacementRepository)
	doctorUsecase := usecase.NewDoctorUsecase(doctorRepository, doctorScheduleRepository, appointmentRepository, redisRepository)
	medicineUsecase := usecase.NewMedicineUsecase(medicineRepository)
	clinicInventoryUsecase := usecase.NewClinicInventoryUsecase(clinicInventoryRepository)
	appointmentUsecase := usecase.NewAppointmentUseCase(db, appointmentRepository, patientRepository, billingRepository, paymentRepository, doctorScheduleRepository, redisRepository, paymentUsecase)
	medicalRecordUsecase := usecase.NewMedicalRecordUsecase(medicalRecordRepository, doctorRepository, redisRepository)
	prescriptionUsecase := usecase.NewPrescriptionUsecase(prescriptionRepository, prescriptionItemRepository, doctorRepository, redisRepository)

	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	patientHandler := handler.NewPatientHandler(patientUsecase)
	clinicHandler := handler.NewClinicHandler(clinicUsecase)
	doctorHandler := handler.NewDoctorHandler(doctorUsecase)
	medicineHandler := handler.NewMedicineHandler(medicineUsecase)
	clinicInventoryHandler := handler.NewClinicInventoryHandler(clinicInventoryUsecase)
	appointmentHandler := handler.NewAppointmentHandler(appointmentUsecase)
	medicalRecordHandler := handler.NewMedicalRecordHandler(medicalRecordUsecase)
	prescriptionHandler := handler.NewPrescriptionHandler(prescriptionUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)

	http.UserRoute(server, userHandler)
	http.AuthRoute(server, authHandler)
	http.PatientRoute(server, patientHandler)
	http.ClinicRoute(server, clinicHandler, clinicInventoryHandler)
	http.DoctorRoute(server, doctorHandler)
	http.MedicineRoute(server, medicineHandler)
	http.AppointmentRoute(server, appointmentHandler)
	http.MedicalRecordRoute(server, medicalRecordHandler)
	http.PrescriptionRoute(server, prescriptionHandler)
	http.PaymentRoute(server, paymentHandler)

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
