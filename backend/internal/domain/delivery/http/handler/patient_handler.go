package handler

import "github.com/gin-gonic/gin"

type PatientHandler interface {
	Me(c *gin.Context)
	Update(c *gin.Context)
}
