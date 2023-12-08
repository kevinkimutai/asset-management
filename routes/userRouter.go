package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/users", handler.Login)
	api.Get("/user/:userId", handler.SignUp)
}
