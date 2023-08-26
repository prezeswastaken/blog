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

// This is the new function to get all posts
func PostGetAll(c *fiber.Ctx) error {
	var posts []models.Post                     // declare a slice of models.Post structs
	initializers.DB.Find(&posts)                // query all records from posts table
	return c.Status(fiber.StatusOK).JSON(posts) // send 200 OK response with posts slice as JSON body
}
