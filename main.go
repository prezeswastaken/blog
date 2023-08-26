package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
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

	// Routes
	app.Get("/", controllers.PostIndex)

	app.Post("/create", controllers.PostCreate)

	app.Listen(":" + os.Getenv("PORT"))
}
