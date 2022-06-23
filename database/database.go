package database

import (
	"fmt"
	"practice/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Migration() {

	DSN := ("host=localhost user=postgres password=sparkman13 dbname=banking_system port=5432 sslmode=disable")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")
	DB.AutoMigrate(&model.User{}, &model.Cars{}, &model.Transaction{}, &model.Message{})
}
