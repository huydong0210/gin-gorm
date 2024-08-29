package helper

import (
	"fmt"
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

func GenerateToken(user *model.User, roles []model.Role, key string) (string, error) {
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

	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

func ParseToken(tokenString string, key string) (*jwt.Token, error) {
	secretKey := []byte(key)
	var result CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &result, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	return token, err
}
