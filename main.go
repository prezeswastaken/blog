package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prezeswastaken/blog/controllers"
	"github.com/prezeswastaken/blog/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDoDatabase()
	initializers.SyncDb()
}

func main() {
	// Setup app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes
	app.Get("/api/get-all", controllers.PostGetAll)

	app.Post("api/create", controllers.PostCreate)

	app.Delete("api/delete", controllers.PostDelete)

	// Serve api
	app.Listen(":" + os.Getenv("PORT"))
}
