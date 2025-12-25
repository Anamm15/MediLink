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

	server := gin.Default()
	server.Use(middlewares.SetupCORS())

	db := database.SetUpDatabaseConnection()
	redisClient := database.NewRedisClient()

	userRepository := repository.NewUserRepository(db)
	patientRepository := repository.NewPatientRepository(db)
	redisRepository := repository.NewRedisRepository(redisClient)
	clinicRepository := repository.NewClinicRepository(db)
	doctorRepository := repository.NewDoctorRepository(db)
	doctorClinicReplacementRepository := repository.NewDoctorClinicPlacementRepository(db)
	doctorScheduleRepository := repository.NewDoctorScheduleRepository(db)
	medicineRepository := repository.NewMedicineRepository(db)
	appointmentRepository := repository.NewAppointmentRepository(db)
	prescriptionRepository := repository.NewPrescriptionRepository(db)
	prescriptionItemRepository := repository.NewPrescriptionItemRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, patientRepository, redisRepository)
	patientUsecase := usecase.NewPatientUsecase(patientRepository)
	clinicUsecase := usecase.NewClinicUsecase(clinicRepository, doctorClinicReplacementRepository)
	doctorUsecase := usecase.NewDoctorUsecase(doctorRepository, doctorScheduleRepository)
	medicineUsecase := usecase.NewMedicineUsecase(medicineRepository)
	appointmentUsecase := usecase.NewAppointmentUseCase(appointmentRepository)
	prescriptionUsecase := usecase.NewPrescriptionUsecase(prescriptionRepository, prescriptionItemRepository)

	userHandler := handler.NewUserHandler(userUsecase)
	patientHandler := handler.NewPatientHandler(patientUsecase)
	clinicHandler := handler.NewClinicHandler(clinicUsecase)
	doctorHandler := handler.NewDoctorHandler(doctorUsecase)
	medicineHandler := handler.NewMedicineHandler(medicineUsecase)
	appointmentHandler := handler.NewAppointmentHandler(appointmentUsecase)
	prescriptionHandler := handler.NewPrescriptionHandler(prescriptionUsecase)

	http.UserRoute(server, userHandler)
	http.PatientRoute(server, patientHandler)
	http.ClinicRoute(server, clinicHandler)
	http.DoctorRoute(server, doctorHandler)
	http.MedicineRoute(server, medicineHandler)
	http.AppointmentRoute(server, appointmentHandler)
	http.PrescriptionRoute(server, prescriptionHandler)

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
