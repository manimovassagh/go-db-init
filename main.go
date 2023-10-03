package main

import (
	databaseinit "github.com/manimovassagh/gorm/databaseInit"
	"github.com/manimovassagh/gorm/models"
)

func init() {

	databaseinit.DBInit()
}
func main() {
	p := models.Product{Code: "Is really fun 2", Price: 66}
	databaseinit.Repository.Create(&p)
}
