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
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)

	routes.UserRoute(server, userDelivery)

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
