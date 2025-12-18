package routes

import (
	"MediLink/internal/domain/delivery"
	"MediLink/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, userDelivery delivery.UserDelivery) {
	user := server.Group("/api/v1/users")
	{
		user.GET("/", middlewares.Authenticate(), middlewares.AuthorizeRole("admin"), userDelivery.GetAll)
		user.GET("/me", middlewares.Authenticate(), userDelivery.GetProfile)

		user.POST("/register", userDelivery.Register)
		user.POST("/login", userDelivery.Login)
		user.POST("/send-otp", middlewares.Authenticate(), userDelivery.SendOTP)
		user.POST("/verify-otp", middlewares.Authenticate(), userDelivery.VerifyOTP)
		user.POST("/on-board-patient", middlewares.Authenticate(), userDelivery.OnBoardPatient)

		user.PUT("/", middlewares.Authenticate(), userDelivery.UpdateProfile)
		user.PATCH("/password", middlewares.Authenticate(), userDelivery.ChangePassword)

		user.DELETE("/", middlewares.Authenticate(), userDelivery.Delete)
	}
}
