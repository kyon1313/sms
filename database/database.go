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

	DSN := ("host=ec2-34-200-35-222.compute-1.amazonaws.com	user=gyiwbhsoxedqid password=8c2e343479959e49e63e6ee94040102f87fae90af23534f2377f87974b0865ed dbname=df6d0e1h8vfbpl port=5432")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")
	DB.AutoMigrate(&model.User{}, &model.Cars{}, &model.Transaction{}, &model.Message{})
}
