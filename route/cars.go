package route

import (
	"practice/database"
	"practice/model"

	"github.com/gofiber/fiber/v2"
)

func AddCars(c *fiber.Ctx) error {
	var cars model.Cars

	if err := c.BodyParser(&cars); err != nil {
		c.JSON(&fiber.Map{
			"sucess": false,
		})
	}
	database.DB.Create(&cars)

	return c.JSON(&fiber.Map{
		"message": "cars added",
		"sucesss": true,
		"data":    cars,
	})
}

func GetCars(c *fiber.Ctx) error {
	var car []model.Cars

	database.DB.Find(&car)
	if len(car) == 0 {
		return c.JSON(&fiber.Map{
			"message": "Cars Fetched failed",
			"success": false,
		})
	}
	return c.JSON(&fiber.Map{
		"message": "Cars successfully Fetched",
		"data":    car,
		"success": true,
	})

}
