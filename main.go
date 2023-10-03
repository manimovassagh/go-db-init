package main

import (
	databaseinit "github.com/manimovassagh/gorm/databaseInit"
	"github.com/manimovassagh/gorm/models"
)

func init() {

	databaseinit.DBInit()
}
func main() {
	//create a list of products for me
	products := []models.Product{
		{Code: "Is really fun", Price: 66},
		{Code: "Is really fun 2", Price: 66},
		{Code: "Is really fun 3", Price: 66},
		{Code: "Is really fun 3", Price: 66},
		{Code: "Is really fun 3", Price: 66},
		{Code: "Is really fun 3", Price: 66},
		{Code: "Is really fun 3", Price: 66},
	}
	databaseinit.Repository.Create(&products)

}
