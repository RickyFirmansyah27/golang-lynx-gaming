package services

import (
	"log"

	"go-fiber-vercel/config"
	"go-fiber-vercel/models"

	"github.com/gofiber/fiber/v2"
)

func GetSkins(c *fiber.Ctx) (int, []models.Skins, error) {
	log.Println("[SkinsService] - Fetching items...", c.Queries())

	queryParams := make(map[string]string)
	c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryParams[string(key)] = string(value)
	})

	skins, totalData, err := config.GetAllskins(queryParams)
	if err != nil {
		log.Printf("[SkinsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	log.Printf("[SkinsService] - Successfully fetched %d items", len(skins))
	return totalData, skins, nil
}
