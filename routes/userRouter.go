package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllUsers)
	api.Get("/:userId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetUserById)
	api.Patch("/:userId", handler.Protected, handler.Restricted("superadmin"), handler.UpdateUser)
	api.Delete("/:userId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteUser)

}
