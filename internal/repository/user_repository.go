package repository

import (
	"gorm.io/gorm"
	model "todo-list-gin-gorm/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (repo *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := repo.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
func (repo *UserRepository) FindUserById(id int64) (*model.User, error) {
	var user model.User
	result := repo.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}
func (repo *UserRepository) CreateUser(user *model.User) error {
	return repo.db.Create(user).Error
}
func (repo *UserRepository) FindAllUser() (users []model.User, err error) {
	err = repo.db.Find(&users).Error
	return
}
