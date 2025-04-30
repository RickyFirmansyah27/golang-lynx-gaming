package routes

import (
	"go-fiber-vercel/helpers"

	"github.com/gofiber/fiber/v2"
)

func RootRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return helpers.Success(c, "Welcome to Go Fiber Vercel", nil)
	})
}
