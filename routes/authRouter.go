package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(api fiber.Router) {
	api.Post("/login", handler.Login)
	api.Post("/signup", handler.SignUp)
}
