package service

import (
	model "todo-list-gin-gorm/internal/models"
	"todo-list-gin-gorm/internal/repository"
)

type UserServiceInterface interface {
	FindUserByUserName(username string) (*model.User, error)
	CreateUser(user *model.User, roleId uint) error
	FindAllUsers() ([]model.User, error)
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
func (s *UserService) CreateUser(user *model.User, roleId uint) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	err = s.repo.InsertUserRole(user.ID, roleId)
	return err
}
func (s *UserService) FindAllUsers() ([]model.User, error) {
	return s.repo.FindAllUser()
}
