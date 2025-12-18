package routes

import (
	"MediLink/internal/domain/delivery"
	"MediLink/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func PatientRoute(server *gin.Engine, patientDelivery delivery.PatientDelivery) {
	patient := server.Group("/api/v1/patients")
	{
		patient.PUT("/", middlewares.Authenticate(), patientDelivery.Update)
	}
}
