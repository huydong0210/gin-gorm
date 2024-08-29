package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	request2 "todo-list-gin-gorm/internal/api/request"
	"todo-list-gin-gorm/internal/middleware"
	model "todo-list-gin-gorm/internal/models"
	"todo-list-gin-gorm/internal/service"
)

type TodoItemHandler struct {
	TodoItemService service.TodoItemServiceInterface
	UserService     service.UserServiceInterface
}

func NewTodoItemHandler(todoItemService service.TodoItemServiceInterface, userService service.UserServiceInterface) *TodoItemHandler {
	return &TodoItemHandler{TodoItemService: todoItemService, UserService: userService}
}
func (h *TodoItemHandler) CreateTodoItem(c *gin.Context) {
	var request request2.TodoItemCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, oke := c.Get(middleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(middleware.UserPrincipal)

	user, err := h.UserService.FindUserByUserName(userPrincipal.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.TodoItemService.CreateTodoItem(&model.TodoItem{
		Name:   request.Name,
		State:  request.State,
		UserId: user.ID,
	})
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
func (h *TodoItemHandler) UpdateTodoItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var request request2.TodoItemUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, oke := c.Get(middleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(middleware.UserPrincipal)
	user, err := h.UserService.FindUserByUserName(userPrincipal.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item := &model.TodoItem{
		Name:  request.Name,
		State: request.State,
	}
	err = h.TodoItemService.UpdateTodoItem(uint(id), user.ID, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})

}
func (h *TodoItemHandler) DeleteTodoItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	value, oke := c.Get(middleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(middleware.UserPrincipal)
	user, err := h.UserService.FindUserByUserName(userPrincipal.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.TodoItemService.DeleteTodoItem(uint(id), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})

}
func (h *TodoItemHandler) FindTodoItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	value, oke := c.Get(middleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(middleware.UserPrincipal)
	user, err := h.UserService.FindUserByUserName(userPrincipal.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	item, err := h.TodoItemService.FindTodoItemById(uint(id), user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}
