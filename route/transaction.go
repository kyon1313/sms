package route

import (
	"practice/database"
	"practice/model"

	"github.com/gofiber/fiber/v2"
)

type TransactionS struct {
	ID   uint       `json:"id"`
	User model.User `json:"user"`
	Car  model.Cars `json:"car"`
}

func TransactionSerializer(trans model.Transaction, car model.Cars, user model.User) TransactionS {
	return TransactionS{
		ID:   trans.ID,
		User: user,
		Car:  car,
	}
}

func AddTransaction(c *fiber.Ctx) error {
	var transac model.Transaction
	var user model.User
	var cars model.Cars

	if err := c.BodyParser(&transac); err != nil {
		return c.SendString("transation failed")
	}
	database.DB.Find(&user, "id=?", transac.UserID)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"message": "user not exist",
			"messags": "transaction failed",
		})
	}
	database.DB.Find(&cars, "id=?", transac.CarID)
	if cars.ID == 0 {
		return c.JSON(&fiber.Map{
			"message": "cars not exist",
			"messags": "transaction failed",
		})
	}
	database.DB.Create(&transac)
	database.DB.First(&cars, transac.CarID)
	database.DB.First(&user, transac.UserID)
	respons := TransactionSerializer(transac, cars, user)
	return c.JSON(&fiber.Map{
		"message": "transaction success",
		"success": true,
		"data":    respons,
	})

	//not done im bored
}
