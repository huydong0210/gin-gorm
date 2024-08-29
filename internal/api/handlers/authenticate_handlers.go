package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo-list-gin-gorm/internal/helper"
	model "todo-list-gin-gorm/internal/models"
	"todo-list-gin-gorm/internal/service"
)

type AuthenticateHandlers struct {
	userService service.UserServiceInterface
	roleService service.RoleServiceInterface
}

func NewAuthenticateHandlers(userService service.UserServiceInterface, roleService service.RoleServiceInterface) *AuthenticateHandlers {
	return &AuthenticateHandlers{userService: userService, roleService: roleService}
}
func (h *AuthenticateHandlers) SignIn(c *gin.Context) {
	type loginInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userService.FindUserByUserName(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
		return
	}
	if !helper.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	roles, err := h.roleService.FindAllRolesByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	token, err := helper.GenerateToken(user, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
func (h *AuthenticateHandlers) SignUp(c *gin.Context) {
	type signUpInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	var input signUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userService.FindUserByUserName(input.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}
	hashPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	user = &model.User{
		Username: input.Username,
		Password: hashPassword,
		Email:    input.Email,
	}
	if err := h.userService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sign up successfully"})

}
