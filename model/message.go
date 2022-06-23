package model

type Message struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserId     int    `json:"user_id"`
	User       User   `gorm:"foreignKey:UserId"`
	TheMessage string `json:"the_message"`
	TheCode    int    `json:"the_code"`
}
