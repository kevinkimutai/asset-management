package main

import (
	"asset-management/database"
	router "asset-management/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	//Initialise Database
	database.InitMigration()

	app := fiber.New()
	//Logger middleware
	app.Use(logger.New())
	//ENV variables middleware
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//API routes

	app.Route("/api/v1/auth", router.AuthRouter)
	app.Route("/api/v1/user", router.UserRouter)
	app.Route("/api/v1/asset", router.AssetRouter)
	app.Route("/api/v1/assetType", router.AssetTypeRouter)
	app.Route("/api/v1/condition", router.ConditionRouter)

	app.Listen(":8000")

}
