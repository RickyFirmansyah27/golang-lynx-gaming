package controllers

import (
	"log"

	"go-fiber-vercel/helpers"
	"go-fiber-vercel/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Incoming request with query params: %v", c.Queries())

	totalData, items, err := services.GetItems(c)
	if err != nil {
		log.Printf("[ItemsController] - Failed to fetch items: %v", err)
		return helpers.Error(c, 400, "Failed to fetch items", err)
	}

	data := []any{
		fiber.Map{
			"total_data": totalData,
			"items":      items,
		},
	}

	log.Printf("[ItemsController] - Successfully fetched items")

	return helpers.Success(c, "Successfully fetched items", data)
}
