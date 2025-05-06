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

func CreateItem(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Creating new item")

	item, err := services.CreateItem(c)
	if err != nil {
		log.Printf("[ItemsController] - Failed to create item: %v", err)
		return helpers.Error(c, 400, "Failed to create item", err)
	}

	data := fiber.Map{
		"item": item,
	}

	log.Printf("[ItemsController] - Successfully created item with ID: %d", item.ID)
	return helpers.Success(c, "Successfully created item", data)
}

func UpdateItem(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Updating item with ID: %s", c.Params("id"))

	item, err := services.UpdateItem(c)
	if err != nil {
		log.Printf("[ItemsController] - Failed to update item: %v", err)
		return helpers.Error(c, 400, "Failed to update item", err)
	}

	data := fiber.Map{
		"item": item,
	}

	log.Printf("[ItemsController] - Successfully updated item with ID: %d", item.ID)
	return helpers.Success(c, "Successfully updated item", data)
}

func DeleteItem(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Deleting item with ID: %s", c.Params("id"))

	if err := services.DeleteItem(c); err != nil {
		log.Printf("[ItemsController] - Failed to delete item: %v", err)
		return helpers.Error(c, 400, "Failed to delete item", err)
	}

	log.Printf("[ItemsController] - Successfully deleted item with ID: %s", c.Params("id"))
	return helpers.Success(c, "Successfully deleted item", nil)
}
