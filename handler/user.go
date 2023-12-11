package handler

import (
	"asset-management/database"
	"asset-management/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	//GET ALL USERS AND THE COUNT OF THEIR ASSETS
	users := new([]model.User)

	err := database.DB.Model(&model.User{}).
		Select("users.first_name, users.last_name, users.email, users.designation, COUNT(assets.id) as assets").
		Joins("JOIN assets ON users.id = assets.user_id").
		Group("users.id").
		Find(&users).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&users)

}

func GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	user := new(model.User)
	err := database.DB.Preload("asset").First(&user, userId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such User Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	//ONLY SUPERADMIN CAN DELETE USER
	userId := c.Params("userId")

	user := new(model.User)

	//Check If User Exists
	err := database.DB.First(&user, userId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such User Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Delete(&user, userId).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdateUser(c *fiber.Ctx) error {
	type Body struct {
		role string
	}
	changeRole := new(Body)
	userId := c.Params("userId")
	user := new(model.User)
	userRole := c.BodyParser(&changeRole)

	//Check If User Exists
	err := database.DB.First(&user, userId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such User Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Update("role", userRole).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)

}
