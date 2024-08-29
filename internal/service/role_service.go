package service

import (
	model "todo-list-gin-gorm/internal/models"
	"todo-list-gin-gorm/internal/repository"
)

type RoleServiceInterface interface {
	FindAllRolesByUserId(userId uint) ([]model.Role, error)
	FindRoleByName(roleName string) (model.Role, error)
}
type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repository *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repository}
}
func (s RoleService) FindAllRolesByUserId(userId uint) ([]model.Role, error) {
	return s.repo.FindAllRolesByUserId(userId)
}
func (s RoleService) FindRoleByName(roleName string) (model.Role, error) {
	return s.repo.FindRoleByRoleName(roleName)
}
