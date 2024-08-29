package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const ADMIN = "ADMIN"
const USER = "USER"

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, oke := c.Get(USER_PRICIPAL_CONTEXT_KEY)
		user := value.(UserPrincipal)
		if !oke {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "internal server error",
			})
			c.Abort()
			return
		}
		roles := strings.Split(user.Role, " ")
		hasRole := false
		for _, role := range roles {
			if role == requiredRole {
				hasRole = true
				break
			}
		}
		if !hasRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you don't have role " + requiredRole,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
