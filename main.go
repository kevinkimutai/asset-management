package main

import (
	"asset-management/database"
	router "asset-management/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	//Initialise Database
	database.InitMigration()

	app := fiber.New()
	//Logger middleware
	app.Use(logger.New())

	//API routes
	// app.Group("/api/v1/auth", router.AuthRouter)

	app.Route("/api/v1/auth", router.AuthRouter)
	app.Route("/api/v1/user", router.UserRouter)

	app.Listen(":8000")

}
