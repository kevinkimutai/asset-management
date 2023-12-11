package handler

import (
	"asset-management/database"
	"asset-management/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetAllUsers(c *fiber.Ctx) error {
	//GET ALL USERS AND THE COUNT OF THEIR ASSETS
	users := new([]model.User)

	// err := database.DB.Model(&model.User{}).
	// 	Select("users.first_name, users.last_name, users.email, users.designation, COUNT(assets.id) as assets").
	// 	Joins("JOIN assets ON users.id = assets.user_id").
	// 	Group("users.id").
	// 	Find(&users).Error

	err := database.DB.Preload("Asset").Omit("password").Find(&users).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&users)

}

func GetUserById(c *fiber.Ctx) error {
	//TODO:HANDLE NORECORD ERRORS BETTER
	userId := c.Params("userId")

	user := new(model.User)
	err := database.DB.Preload("Asset").Omit("password").First(&user, userId).Error

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
	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateUser(c *fiber.Ctx) error {
	type Body struct {
		Role string `json:"role"`
	}
	changeRole := new(Body)
	userId := c.Params("userId")
	user := new(model.User)

	if err := c.BodyParser(&changeRole); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	log.Info(&changeRole)

	//Check If User Exists
	err := database.DB.First(&user, userId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such User Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Model(&model.User{}).Where("id = ?", userId).Update("role", &changeRole.Role).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)

}
