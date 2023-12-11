package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func ConditionRouter(api fiber.Router) {
	api.Post("/", handler.Protected, handler.Restricted("superadmin"), handler.CreateCondition)
	api.Get("/", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllConditions)
	api.Get("/:conditionId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetConditionById)
	api.Patch("/:conditionId", handler.Protected, handler.Restricted("superadmin,admin"), handler.UpdateCondition)
	api.Delete("/:conditionId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteCondition)

}
