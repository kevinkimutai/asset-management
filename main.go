package main

import (
	"asset-management/database"
	router "asset-management/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

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

	//CORS Middleware
	//Default CORS middleware
	//app.Use(cors.New())

	//ONly Use In DevelopmenT
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	//USE in Production
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "https://gofiber.io, https://gofiber.net",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	//Compress Middleware
	app.Use(compress.New())

	//CSRF MIddleware
	app.Use(csrf.New())

	//Protect HTTP Headers
	app.Use(helmet.New())

	//API Limiter
	app.Use(limiter.New())

	//API routes
	app.Route("/api/v1/auth", router.AuthRouter)
	app.Route("/api/v1/users", router.UserRouter)
	app.Route("/api/v1/asset", router.AssetRouter)
	app.Route("/api/v1/assetType", router.AssetTypeRouter)
	app.Route("/api/v1/condition", router.ConditionRouter)

	//Monitor API Metrics
	app.Get("/metrics", monitor.New())

	PORT := os.Getenv("PORT")

	app.Listen(":" + PORT)

}
