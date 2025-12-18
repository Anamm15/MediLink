package delivery

import "github.com/gin-gonic/gin"

type PatientDelivery interface {
	Update(c *gin.Context)
}
