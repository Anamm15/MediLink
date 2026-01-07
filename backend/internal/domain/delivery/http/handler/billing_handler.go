package handler

import "github.com/gin-gonic/gin"

type BillingHandler interface {
	GetByPatient(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}
