package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
app.Get("/todos", func(c *fiber.Ctx) error {
	// Logic to fetch all todos from the database
	// and return them as a JSON response
	return c.JSON([]Todo{})
})

app.Post("/todos", func(c *fiber.Ctx) error {
	// Logic to create a new todo based on the request body
	// and save it to the database
	return c.SendString("Todo created successfully")
})

app.Get("/todos/:id", func(c *fiber.Ctx) error {
	// Logic to fetch a specific todo from the database
	// based on the provided ID and return it as a JSON response
	return c.JSON(Todo{})
})

app.Put("/todos/:id", func(c *fiber.Ctx) error {
	// Logic to update a specific todo in the database
	// based on the provided ID and request body
	return c.SendString("Todo updated successfully")
})

app.Delete("/todos/:id", func(c *fiber.Ctx) error {
	// Logic to delete a specific todo from the database
	// based on the provided ID
	return c.SendString("Todo deleted successfully")
})