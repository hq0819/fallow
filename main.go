package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/demo", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")

	})

	app.Listen(":8000")

}
