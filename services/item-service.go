package services

import (
	"go-fiber-vercel/config"
	"go-fiber-vercel/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetSkins(c *fiber.Ctx) (int, []models.Skins, error) {
	log.Println("[SkinsService] - Fetching items...", c.Queries())

	// Get query parameters from Fiber context
	queryParams := make(map[string]string)
	c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryParams[string(key)] = string(value)
	})

	itemsMap, totalData, err := config.GetAllskins(queryParams)
	if err != nil {
		log.Printf("[SkinsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	items := make([]models.Skins, 0, len(itemsMap))
	for _, itemMap := range itemsMap {
		item := models.Skins{
			ID:     itemMap["id"].(int),
			Name:   itemMap["nama"].(string),
			Hero:   itemMap["hero"].(string),
			Tag:    itemMap["tag"].(string),
			Config: itemMap["config"].(byte),
		}
		items = append(items, item)
	}

	log.Printf("[SkinsService] - Successfully fetched %d items", len(items))
	return totalData, items, nil
}
