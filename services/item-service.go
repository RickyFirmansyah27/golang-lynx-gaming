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
