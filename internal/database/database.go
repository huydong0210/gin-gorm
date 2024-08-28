package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	model "todo-list-gin-gorm/internal/models"
)

func Initialize(databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.TodoItem{},
	)
}
