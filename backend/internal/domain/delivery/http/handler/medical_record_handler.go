package handler

import "github.com/gin-gonic/gin"

type MedicalRecordHandler interface {
	GetByPatient(c *gin.Context)
	GetByDoctor(c *gin.Context)
	GetDetailByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
