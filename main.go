package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prezeswastaken/blog/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDoDatabase()
}

// Define a struct for the JSON response
type Message struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := Message{
			Name: "Kacper",
			Age:  21,
		}
		return c.JSON(data)
	})

	app.Listen(":3000")
}
