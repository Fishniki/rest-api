package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/developers", develoPers)
	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}

func develoPers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"name" : "Niki",
		"kelas" : "XII RPL",
		"Hobi" : "Coding",
	})
}