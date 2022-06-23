package route

import (
	"practice/database"
	"practice/model"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		c.JSON(&fiber.Map{
			"message": "adding user failed",
			"data":    user,
		})
	}

	database.DB.Create(&user)
	return c.JSON(&fiber.Map{
		"message": "user added successdfully",
		"success": true,
		"data":    user,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var user []model.User
	database.DB.Find(&user)
	if len(user) == 0 {
		return c.SendString("fukc youn no user found")
	}
	return c.JSON(&fiber.Map{
		"message": "user fetch successdfully",
		"success": true,
		"data":    user,
	})
}

func GetUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User

	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"message": "user not exist",
		})
	}
	return c.JSON(&fiber.Map{
		"message": "user fetch successdfully",
		"success": true,
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"message": "Update Failed",
			"success": false,
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(&fiber.Map{
			"message": "Update Failed",
			"error":   err,
		})
	}
	database.DB.Save(&user)
	return c.JSON(&fiber.Map{
		"message": "Update Sucess",
		"data":    user,
		"success": true,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"message": "User Not Exist",
			"success": false,
		})
	}

	database.DB.Delete(&user)
	return c.JSON(&fiber.Map{
		"message": "User  Sucessfully deleted",
		"data":    user,
		"success": true,
	})

}
