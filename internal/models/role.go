package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null; unique"`
}

func FindRolesByUserId(db *gorm.DB, userId uint) ([]Role, error) {
	var roles []Role
	result := db.Raw("select * from roles where id in ( select role_id from user_role where user_id = ?)", userId).Scan(&roles)
	return roles, result.Error
}
func FindRoleByRoleName(db *gorm.DB, roleName string) (Role, error) {
	var role Role
	result := db.Raw("select * from roles where name = ?", roleName).Scan(&role)
	return role, result.Error
}
