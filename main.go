package main

import (
	"github.com/manimovassagh/gorm/db"
	"github.com/manimovassagh/gorm/models"
)

func init() {

	db.DBInit()
}
func main() {
	//create a list of 10 sample products for me

	products := []models.Product{
		{Code: "P1", Price: 1000, Case: "Electronics", Amount: 50},
		{Code: "P2", Price: 500, Case: "Electronics", Amount: 100},
		{Code: "P3", Price: 150, Case: "Furniture", Amount: 30},
		{Code: "P4", Price: 800, Case: "Electronics", Amount: 20},
		{Code: "P5", Price: 200, Case: "Furniture", Amount: 40},
		{Code: "P6", Price: 80, Case: "Apparel", Amount: 60},
		{Code: "P7", Price: 50, Case: "Appliances", Amount: 25},
		{Code: "P8", Price: 300, Case: "Furniture", Amount: 15},
		{Code: "P9", Price: 70, Case: "Outdoor Gear", Amount: 35},
		{Code: "P10", Price: 400, Case: "Electronics", Amount: 10},
	}

	db.Repository.Create(&products)

}
