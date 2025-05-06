package routes

import (
	"go-fiber-vercel/controllers"

	"github.com/gofiber/fiber/v2"
)

// V1Route mengatur routing untuk versi 1
func V1Route(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/items", controllers.GetAllItems)
	v1.Post("/items", controllers.CreateItem)
	v1.Patch("/items/:id", controllers.UpdateItem)
	v1.Delete("/items/:id", controllers.DeleteItem)
}
