package repository

import (
	"gorm.io/gorm"
	model "todo-list-gin-gorm/internal/models"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}
func (r *RoleRepository) FindAllRolesByUserId(userId uint) ([]model.Role, error) {
	var roles []model.Role
	result := r.db.Raw("select * from roles where id in ( select role_id from user_role where user_id = ?)", userId).Scan(&roles)
	return roles, result.Error
}
func (r *RoleRepository) FindRoleByRoleName(roleName string) (model.Role, error) {
	var role model.Role
	result := r.db.Raw("select * from roles where name = ?", roleName).Scan(&role)
	return role, result.Error
}
