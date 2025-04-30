package main

import (
	"go-fiber-vercel/config"
	"go-fiber-vercel/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	err := config.DBConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	routes.RootRoute(app)
	routes.V1Route(app)
	routes.V2Route(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("[Fiber-Service] Server is running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
