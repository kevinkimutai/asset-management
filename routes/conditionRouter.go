package router

import (
	"asset-management/handler"

	"github.com/gofiber/fiber/v2"
)

func ConditionRouter(api fiber.Router) {
	api.Post("/condition", handler.Protected, handler.Restricted("superadmin"), handler.CreateCondition)
	api.Get("/conditions", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetAllConditions)
	api.Get("/condition/:conditionId", handler.Protected, handler.Restricted("superadmin", "admin"), handler.GetConditionById)
	api.Patch("/condition/:conditionId", handler.Protected, handler.Restricted("superadmin,admin"), handler.UpdateCondition)
	api.Delete("/condition/:conditionId", handler.Protected, handler.Restricted("superadmin"), handler.DeleteCondition)

}
