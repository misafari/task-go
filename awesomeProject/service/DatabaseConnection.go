package service

import (
	"../model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := "admin:admin@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.Task{})
	database.AutoMigrate(&model.User{})

	DB = database
	//database.Create(&Product{Code: "D42", Price: 100})
	//
	//var product Product
	//database.First(&product, 1) // find product with integer primary key
	//database.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//database.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//database.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	//// non-zero fields
	//database.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//database.Delete(&product, 1)
}