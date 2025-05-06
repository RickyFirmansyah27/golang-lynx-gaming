package services

import (
	"go-fiber-vercel/config"
	"go-fiber-vercel/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) (int, []models.Item, error) {
	log.Println("[ItemsService] - Fetching items...", c.Queries())

	// Get query parameters from Fiber context
	queryParams := make(map[string]string)
	c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryParams[string(key)] = string(value)
	})

	itemsMap, totalData, err := config.GetAllItems(queryParams)
	if err != nil {
		log.Printf("[ItemsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	items := make([]models.Item, 0, len(itemsMap))
	for _, itemMap := range itemsMap {
		item := models.Item{
			ID:         itemMap["id"].(int),
			Name:       itemMap["name"].(string),
			CategoryID: itemMap["category_id"].(int),
			Stock:      itemMap["stock"].(int),
			Unit:       itemMap["unit"].(string),
			MinStock:   itemMap["min_stock"].(int),
		}
		items = append(items, item)
	}

	log.Printf("[ItemsService] - Successfully fetched %d items", len(items))
	return totalData, items, nil
}

func CreateItem(c *fiber.Ctx) (*models.Item, error) {
	log.Println("[ItemsService] - Creating new item")

	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		log.Printf("[ItemsService] - Error parsing request body: %v", err)
		return nil, err
	}

	createdItem, err := config.CreateItem(item)
	if err != nil {
		log.Printf("[ItemsService] - Error creating item: %v", err)
		return nil, err
	}

	log.Printf("[ItemsService] - Successfully created item with ID: %d", createdItem.ID)
	return createdItem, nil
}

func UpdateItem(c *fiber.Ctx) (*models.Item, error) {
	log.Println("[ItemsService] - Updating item")

	itemID, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("[ItemsService] - Error parsing item ID: %v", err)
		return nil, err
	}

	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		log.Printf("[ItemsService] - Error parsing request body: %v", err)
		return nil, err
	}

	updatedItem, err := config.UpdateItem(itemID, item)
	if err != nil {
		log.Printf("[ItemsService] - Error updating item: %v", err)
		return nil, err
	}

	log.Printf("[ItemsService] - Successfully updated item with ID: %d", updatedItem.ID)
	return updatedItem, nil
}

func DeleteItem(c *fiber.Ctx) error {
	log.Println("[ItemsService] - Deleting item")

	itemID, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("[ItemsService] - Error parsing item ID: %v", err)
		return err
	}

	if err := config.DeleteItem(itemID); err != nil {
		log.Printf("[ItemsService] - Error deleting item: %v", err)
		return err
	}

	log.Printf("[ItemsService] - Successfully deleted item with ID: %d", itemID)
	return nil
}
