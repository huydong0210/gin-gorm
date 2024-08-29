package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"todo-list-gin-gorm/internal/helper"
)

const USER_PRICIPAL_CONTEXT_KEY = "USER_PRINCIPAL"

type UserPrincipal struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

func JwtMiddleWare(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}
		token, err := helper.ParseToken(tokenString[1], key)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userPrincipal := UserPrincipal{
				Username: claims["username"].(string),
				Role:     claims["role"].(string),
				Email:    claims["email"].(string),
			}
			c.Set(USER_PRICIPAL_CONTEXT_KEY, userPrincipal)
		}
		c.Next()
	}
}
