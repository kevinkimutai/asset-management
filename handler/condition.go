package handler

import (
	"asset-management/database"
	"asset-management/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllConditions(c *fiber.Ctx) error {

	conditions := new([]model.Condition)

	err := database.DB.Find(&conditions).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&conditions)

}

func CreateCondition(c *fiber.Ctx) error {
	condition := new(model.Condition)

	if err := c.BodyParser(&condition); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err := database.DB.Create(&condition).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&condition)

}

func UpdateCondition(c *fiber.Ctx) error {
	condition := new(model.Condition)
	conditionId := c.Params("conditionId")

	if err := c.BodyParser(&condition); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	//Check If Asset Exists
	err := database.DB.First(&model.Condition{}, conditionId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such condition Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Model(&model.Condition{}).Where("id = ?", conditionId).Updates(&condition).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetConditionById(c *fiber.Ctx) error {
	condition := new(model.Condition)
	conditionId := c.Params("assetId")

	//Check If Asset Exists
	err := database.DB.First(&condition, conditionId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such Asset Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&condition)
}

func DeleteCondition(c *fiber.Ctx) error {
	//ONLY SUPERADMIN CAN DELETE ASSET
	conditionId := c.Params("conditionId")

	condition := new(model.Condition)

	//Check If Asset Exists
	err := database.DB.First(&condition, conditionId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such condition Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Delete(&condition, conditionId).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
