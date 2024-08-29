package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo-list-gin-gorm/internal/api/handlers"
	"todo-list-gin-gorm/internal/config"
	"todo-list-gin-gorm/internal/middleware"
	"todo-list-gin-gorm/internal/repository"
	"todo-list-gin-gorm/internal/service"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB, config *config.Config) {

	jwtMiddleware := middleware.JwtMiddleWare(config.SecretKey)
	requireAdmin := middleware.RequireRole(middleware.ADMIN)
	requireUser := middleware.RequireRole(middleware.USER)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	roleRepo := repository.NewRoleRepository(db)
	rolService := service.NewRoleService(roleRepo)

	authService := service.NewAuthenticateService(config, userService, rolService)

	userHandler := handlers.NewUserHandler(userService)

	_ = userHandler

	authenticateHandlers := handlers.NewAuthenticateHandlers(authService)

	authenticateRoutes := router.Group("/api")
	{
		authenticateRoutes.POST("/login", authenticateHandlers.SignIn)
		authenticateRoutes.POST("/sign-up", authenticateHandlers.SignUp)
	}

	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(jwtMiddleware)
	adminRoutes.Use(requireAdmin)
	{
		adminRoutes.GET("/list-users", userHandler.FindAllUsers)
	}

	todoRepo := repository.NewTodoItemRepository(db)
	todoService := service.NewTodoItemService(todoRepo)
	todoHandler := handlers.NewTodoItemHandler(todoService, userService)

	todoItemRoutes := router.Group("/api/todo-item")
	todoItemRoutes.Use(jwtMiddleware)
	todoItemRoutes.Use(requireUser)
	{
		todoItemRoutes.GET("/:id", todoHandler.FindTodoItem)
		todoItemRoutes.POST("", todoHandler.CreateTodoItem)
		todoItemRoutes.DELETE("/:id", todoHandler.DeleteTodoItem)
		todoItemRoutes.PUT("/:id", todoHandler.UpdateTodoItem)
	}

}
