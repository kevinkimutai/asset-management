package handler

import (
	"asset-management/database"
	"asset-management/model"

	"github.com/gofiber/fiber/v2"
)

func CreateAssetType(c *fiber.Ctx) error {
	assetType := new(model.AssetType)

	if err := c.BodyParser(&assetType); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err := database.DB.Create(&assetType).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&assetType)

}
func GetAllAssetTypes(c *fiber.Ctx) error {

	assetTypes := new([]model.AssetType)

	err := database.DB.Find(&assetTypes).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&assetTypes)
}
func GetAssetTypeById(c *fiber.Ctx) error {
	assetTypeId := c.Params("assetTypeId")
	assetType := new(model.Asset)

	err := database.DB.First(&assetType, assetTypeId).Error

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such AssetType With that Id",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&assetType)
}
func UpdateAssetType(c *fiber.Ctx) error {
	assetTypeId := c.Params("assetTypeId")
	assetType := new(model.Asset)

	if err := c.BodyParser(&assetType); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error()})
	}

	//Check if AssetType Exists
	err := database.DB.First(&model.AssetType{}, assetTypeId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such AssetType Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Model(&model.AssetType{}).Where("id = ?", assetType).Updates(&assetType).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
func DeleteAssetType(c *fiber.Ctx) error {
	//ONLY SUPERADMIN CAN DELETE ASSETType
	assetTypeId := c.Params("assetTypeId")

	assetType := new(model.AssetType)

	//Check If Asset Exists
	err := database.DB.First(&assetType, assetTypeId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such Asset Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Delete(&assetType, assetTypeId).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
