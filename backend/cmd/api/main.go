package main

import (
	"log"
	"os"

	"MediLink/internal/delivery"
	"MediLink/internal/infrastructure/database"
	"MediLink/internal/infrastructure/repository"
	"MediLink/internal/middlewares"
	"MediLink/internal/routes"
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

	userUsecase := usecase.NewUserUsecase(userRepository, patientRepository, redisRepository)
	patientUsecase := usecase.NewPatientUsecase(patientRepository)

	userDelivery := delivery.NewUserDelivery(userUsecase)
	patientDelivery := delivery.NewPatientDelivery(patientUsecase)

	routes.UserRoute(server, userDelivery)
	routes.PatientRoute(server, patientDelivery)

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
