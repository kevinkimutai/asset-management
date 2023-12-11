package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func AssetRouter(api fiber.Router) {
	api.Post("/asset", handler.Protected, handler.Restricted("admin"), handler.CreateAsset)
	api.Get("/assets", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllAssets)
	api.Get("/asset/:assetId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAssetById)
	api.Patch("/asset/:assetId", handler.Protected, handler.Restricted("superadmin,admin"), handler.UpdateAsset)
	api.Delete("/asset/:assetId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteAsset)

}
