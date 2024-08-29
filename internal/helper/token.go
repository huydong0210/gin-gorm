package helper

import (
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
	model "todo-list-gin-gorm/internal/models"
)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User, roles []model.Role) (string, error) {
	var roleNames string
	for _, role := range roles {
		roleNames += role.Name + " "
	}
	roleNames = strings.TrimSpace(roleNames)
	claims := CustomClaims{
		Username: user.Username,
		Role:     roleNames,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("dfgdfgdfgd"))
	return tokenString, err
}
