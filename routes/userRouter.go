package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/users", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllUsers)
	api.Get("/user/:userId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetUserById)
	api.Patch("/user/:userId", handler.Protected, handler.Restricted("superadmin"), handler.UpdateUser)
	api.Delete("/user/:userId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteUser)

}
