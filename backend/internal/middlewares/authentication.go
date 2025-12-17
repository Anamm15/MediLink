package middlewares

import (
	"net/http"
	"strings"

	"MediLink/internal/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed("Failed ", "Access token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenStr := strings.Split(authHeader, " ")[1]
		if tokenStr == "" {
			response := utils.BuildResponseFailed("Failed ", "Access token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			response := utils.BuildResponseFailed("Failed ", "Invalid access token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("token", tokenStr)
		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
