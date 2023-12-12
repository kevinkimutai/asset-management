package handler

import (
	"asset-management/database"
	"asset-management/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetAllUsers(c *fiber.Ctx) error {
	//GET ALL USERS AND THE COUNT OF THEIR ASSETS
	users := new([]model.User)
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "20")
	search := c.Query("search", "")

	//convert page&pageSize to Int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page parameter",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid pageSize parameter",
		})
	}

	offset := (pageInt - 1) * pageSizeInt

	query := database.DB.Offset(offset).Limit(pageSizeInt)

	//SearchBy Name
	if search != "" {
		query = query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+search+"%", "%"+search+"%")

	}

	err = query.Preload("Asset").Omit("password").Find(&users).Error

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
