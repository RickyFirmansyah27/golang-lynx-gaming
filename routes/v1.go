package routes

import (
	"go-fiber-vercel/controllers"

	"github.com/gofiber/fiber/v2"
)

// V1Route mengatur routing untuk versi 1
func V1Route(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/login", controllers.Login)
	v1.Post("/register", controllers.Register)
	v1.Post("/check-auth", controllers.CheckAuth)

	v1.Get("/skins", controllers.GetAllSkins)
	v1.Patch("/skins/:id", controllers.UpdateSkins)
	v1.Post("/skins", controllers.CreateSkins)

	v1.Get("/arenas", controllers.GetAllArenas)
	v1.Patch("/arenas/:id", controllers.UpdateArenas)
	v1.Post("/arenas", controllers.CreateArenas)
}
