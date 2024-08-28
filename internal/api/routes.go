package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo-list-gin-gorm/internal/api/handlers"
	"todo-list-gin-gorm/internal/repository"
	"todo-list-gin-gorm/internal/service"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	_ = userHandler

	//example

	//orderRoutes := router.Group("/orders")
	//{
	//	orderRoutes.POST("", orderHandler.CreateOrder)
	//	orderRoutes.GET("/:id", orderHandler.GetOrder)
	//	orderRoutes.GET("", orderHandler.ListOrders)
	//	orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
	//	orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	//}

}
