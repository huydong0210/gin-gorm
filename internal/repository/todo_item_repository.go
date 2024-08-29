package repository

import (
	"gorm.io/gorm"
	model "todo-list-gin-gorm/internal/models"
)

type TodoItemRepository struct {
	db *gorm.DB
}

func NewTodoItemRepository(db *gorm.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (repo TodoItemRepository) CreateToDoItem(item *model.TodoItem) error {
	return repo.db.Create(item).Error
}
func (repo TodoItemRepository) DeleteTodoItem(itemId uint, userId uint) error {
	return repo.db.Where("id = ? and user_id = ?", itemId, userId).Delete(&model.TodoItem{}, itemId).Error
}
func (repo TodoItemRepository) UpdateTodoItem(itemId uint, userId uint, item *model.TodoItem) error {
	return repo.db.Model(&model.TodoItem{}).Where("id = ? and user_id = ?", itemId, userId).Updates(item).Error
}
func (repo TodoItemRepository) FindTodoItemById(itemId uint, userId uint) (*model.TodoItem, error) {
	var item model.TodoItem
	result := repo.db.Where("id = ? and user_id = ?", itemId, userId).First(&item, itemId)
	return &item, result.Error
}
