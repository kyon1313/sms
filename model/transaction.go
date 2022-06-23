package model

type Transaction struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
	CarID  int  `json:"car_id"`
	Cars   Cars `gorm:"foreignKey:CarID"`
}
