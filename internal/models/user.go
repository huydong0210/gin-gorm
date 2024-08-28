package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
}

var AnonymousUser = &User{}

func FindUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}
func DeleteUserByUsername(db *gorm.DB, username string) error {
	result := db.Where("username = ?", username).Delete(&User{})
	return result.Error
}
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}
func InsertUserRoles(db *gorm.DB, userId uint, roleName string) error {
	role, err := FindRoleByRoleName(db, roleName)
	if err != nil {
		return err
	}
	err = db.Exec("insert into user_role(user_id, role_id) value ( ? , ?)", userId, role.ID).Error
	return err
}
func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
func FindAllUsers(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return
}
