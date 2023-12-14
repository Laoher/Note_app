package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectDB() {
	// Connect to DB
	database, err := gorm.Open(mysql.Open("root:B0115588f!@tcp(127.0.0.1:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	log.Println("Database connected!")
	DB = database

}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
}