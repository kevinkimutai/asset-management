package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func AssetRouter(api fiber.Router) {
	api.Post("/", handler.Protected, handler.Restricted("admin"), handler.CreateAsset)
	api.Get("/", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllAssets)
	api.Get("/:assetId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAssetById)
	api.Patch("/:assetId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.UpdateAsset)
	api.Delete("/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteAsset)

}
