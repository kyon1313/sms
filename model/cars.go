package model

type Cars struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	CarName string `json:"car_name"`
	Price   int    `json:"price"`
}
