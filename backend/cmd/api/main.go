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

	userUsecase := usecase.NewUserUsecase(userRepository, patientRepository, redisRepository)
	patientUsecase := usecase.NewPatientUsecase(patientRepository)
	clinicUsecase := usecase.NewClinicUsecase(clinicRepository)

	userHandler := handler.NewUserHandler(userUsecase)
	patientHandler := handler.NewPatientHandler(patientUsecase)
	clinicHandler := handler.NewClinicHandler(clinicUsecase)

	http.UserRoute(server, userHandler)
	http.PatientRoute(server, patientHandler)
	http.ClinicRoute(server, clinicHandler)

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
