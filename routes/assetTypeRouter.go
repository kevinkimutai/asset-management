package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func AssetTypeRouter(api fiber.Router) {
	api.Post("/assetType", handler.Protected, handler.Restricted("admin,superadmin"), handler.CreateAssetType)
	api.Get("/assetType", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllAssetTypes)
	api.Get("/assetType/:assetId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAssetTypeById)
	api.Patch("/assetType/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.UpdateAssetType)
	api.Delete("/assetType/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteAssetType)

}
