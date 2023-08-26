package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prezeswastaken/blog/initializers"
	"github.com/prezeswastaken/blog/models"
)

func PostIndex(c *fiber.Ctx) error {
	return c.SendString("Hello from the postController!")
}

func PostCreate(c *fiber.Ctx) error {
	type PostData struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	postData := new(PostData)

	if err := c.BodyParser(postData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	post := models.Post{
		Title: postData.Title,
		Body:  postData.Body,
	}
	initializers.DB.Create(&post) // pass pointer of data to Create
	return c.SendString("I'm a POST request!")
}
