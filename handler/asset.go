package handler

import (
	"asset-management/database"
	"asset-management/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllAssets(c *fiber.Ctx) error {
	//TODO:HANDLE QUERIES
	assets := new([]model.Asset)
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "20")
	search := c.Query("search", "")

	//convert page&&pageSize tO INt

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

	//SearchBy SetrialNumber
	if search != "" {
		query = query.Where("serial_number LIKE ?", "%"+search+"%")

	}

	err = query.Preload("Condition").Preload("AssetType").Preload("User").Find(&assets).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&assets)

}

func CreateAsset(c *fiber.Ctx) error {
	asset := new(model.Asset)

	if err := c.BodyParser(&asset); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err := database.DB.Create(&asset).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(&asset)

}

func UpdateAsset(c *fiber.Ctx) error {
	asset := new(model.Asset)
	assetId := c.Params("assetId")

	if err := c.BodyParser(&asset); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	//Check If Asset Exists
	err := database.DB.First(&model.Asset{}, assetId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such Asset Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Model(&model.Asset{}).Where("id = ?", assetId).Updates(&asset).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAssetById(c *fiber.Ctx) error {
	asset := new(model.Asset)
	assetId := c.Params("assetId")

	//Check If Asset Exists
	err := database.DB.First(&asset, assetId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such Asset Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&asset)
}

func DeleteAsset(c *fiber.Ctx) error {
	//ONLY SUPERADMIN CAN DELETE ASSET
	assetId := c.Params("assetId")

	asset := new(model.Asset)

	//Check If Asset Exists
	err := database.DB.First(&asset, assetId).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Such Asset Found",
			"error":   err.Error(),
		})
	}

	err = database.DB.Delete(&asset, assetId).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
