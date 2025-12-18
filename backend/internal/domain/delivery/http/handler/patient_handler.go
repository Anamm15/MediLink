package handler

import "github.com/gin-gonic/gin"

type PatientHandler interface {
	Update(c *gin.Context)
}
