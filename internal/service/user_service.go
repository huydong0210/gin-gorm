package service

import (
	model "todo-list-gin-gorm/internal/models"
	"todo-list-gin-gorm/internal/repository"
)

type UserServiceInterface interface {
	FindUserByUserName(username string) (*model.User, error)
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindUserByUserName(username string) (*model.User, error) {
	return s.repo.FindUserByUsername(username)
}