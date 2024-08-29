package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-gin-gorm/internal/service"
)

type AuthenticateHandlers struct {
	authService service.AuthenticateServiceInterface
}

func NewAuthenticateHandlers(authenticateService service.AuthenticateServiceInterface) *AuthenticateHandlers {
	return &AuthenticateHandlers{authService: authenticateService}
}
func (h *AuthenticateHandlers) SignIn(c *gin.Context) {

	var request service.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authService.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
func (h *AuthenticateHandlers) SignUp(c *gin.Context) {
	var request service.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.authService.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sign up successfully"})

}
