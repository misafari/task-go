package service

import (
	"../model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err :=
		gorm.Open(mysql.Open("admin:admin@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"),
			&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := database.AutoMigrate(&model.Task{}); err != nil {
		panic("failed to create task table")
	}

	if err = database.AutoMigrate(&model.User{}); err != nil {
		panic("failed to create user table")
	}

	DB = database
}
