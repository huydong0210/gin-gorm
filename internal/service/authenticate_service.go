package service

import (
	"gorm.io/gorm"
	"todo-list-gin-gorm/internal/config"
	error2 "todo-list-gin-gorm/internal/error"
	"todo-list-gin-gorm/internal/helper"
	model "todo-list-gin-gorm/internal/models"
)

type AuthenticateServiceInterface interface {
	Login(request LoginRequest) (string, error)
	SignUp(request SignUpRequest) error
}
type AuthenticateService struct {
	config      config.Config
	userService UserServiceInterface
	roleService RoleServiceInterface
}
type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthenticateService(config *config.Config, userService UserServiceInterface, roleService RoleServiceInterface) *AuthenticateService {
	return &AuthenticateService{
		config:      *config,
		userService: userService,
		roleService: roleService,
	}
}
func (s *AuthenticateService) Login(request LoginRequest) (string, error) {
	user, err := s.userService.FindUserByUserName(request.Username)
	if err != nil {
		return "", &error2.AppError{
			Message: "username not found",
		}
	}
	if !helper.CheckPasswordHash(request.Password, user.Password) {
		return "", &error2.AppError{
			Message: "wrong password",
		}
	}
	roles, err := s.roleService.FindAllRolesByUserId(user.ID)
	if err != nil {
		return "", &error2.AppError{
			Message: "internal server error",
		}
	}
	token, err := helper.GenerateToken(user, roles, s.config.SecretKey)
	if err != nil {
		return "", &error2.AppError{
			Message: "internal server error",
		}
	}
	return token, nil
}
func (s *AuthenticateService) SignUp(request SignUpRequest) error {

	user, err := s.userService.FindUserByUserName(request.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err != gorm.ErrRecordNotFound {
		return &error2.AppError{
			Message: "username already exists",
		}
	}
	hashPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return &error2.AppError{
			Message: "internal server error",
		}
	}
	user = &model.User{
		Username: request.Username,
		Password: hashPassword,
		Email:    request.Email,
	}
	if err := s.userService.CreateUser(user); err != nil {
		return &error2.AppError{
			Message: "internal server error",
		}
	}
	return nil
}
