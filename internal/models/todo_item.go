package model

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Name   string
	State  string
	UserId uint
}

func CreateTodoItem(db *gorm.DB, item *TodoItem) error {
	result := db.Create(item)
	return result.Error
}
func DeleteTodoItem(db *gorm.DB, id int) error {
	result := db.Delete(&TodoItem{}, id)
	return result.Error
}
func UpdateTodoItem(db *gorm.DB, id int, item *TodoItem) error {
	result := db.Model(&TodoItem{}).Where("id = ?", id).Updates(item)
	return result.Error
}
func FindTodoItemById(db *gorm.DB, id int) (TodoItem, error) {
	var item TodoItem
	result := db.First(&item, id)
	return item, result.Error
}
