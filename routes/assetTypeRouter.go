package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func AssetTypeRouter(api fiber.Router) {
	api.Post("/", handler.Protected, handler.Restricted("admin", "superadmin"), handler.CreateAssetType)
	api.Get("/", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllAssetTypes)
	api.Get("/:assetId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAssetTypeById)
	api.Patch("/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.UpdateAssetType)
	api.Delete("/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteAssetType)

}
