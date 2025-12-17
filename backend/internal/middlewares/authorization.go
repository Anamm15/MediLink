package middlewares

import (
	"net/http"
	"slices"

	"MediLink/internal/helpers/constants"
	"MediLink/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleInterface, exists := c.Get("role")
		if !exists {
			res := utils.BuildResponseFailed("Failed ", "Role not found", nil)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		userRole, ok := roleInterface.(constants.UserRole)
		if !ok {
			roleStr, okStr := roleInterface.(string)
			if !okStr {
				res := utils.BuildResponseFailed("Failed", "Internal Server Error: Invalid Role Type", nil)
				c.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
			userRole = constants.UserRole(roleStr)
		}

		if slices.Contains(allowedRoles, string(userRole)) {
			c.Next()
			return
		}

		res := utils.BuildResponseFailed("Failed ", "Unauthorized", nil)
		c.JSON(http.StatusForbidden, res)
		c.Abort()
	}
}
