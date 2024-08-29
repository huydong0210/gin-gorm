package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-gin-gorm/internal/service"
)

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}
func (h *UserHandler) FindAllUsers(c *gin.Context) {
	users, err := h.service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
