package db

import (
	"fmt"

	"github.com/manimovassagh/gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Repository *gorm.DB

func DBInit() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:8088)/products?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	Repository, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to DB")
		return nil
	}

	fmt.Println("Database connected successfully")
	Repository.Migrator().DropTable(&models.Product{})

	Repository.AutoMigrate(&models.Product{}) // Assuming "models.Product" is defined in your models package

	return Repository
}
