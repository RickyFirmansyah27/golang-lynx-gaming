package services

import (
	"log"
	"strconv"

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

func UpdateSkin(c *fiber.Ctx) (*models.Skins, error) {
	log.Println("[SkinsService] - Updating skin...")

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[SkinsService] - Invalid ID param: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid ID parameter")
	}

	var skin models.Skins
	if err := c.BodyParser(&skin); err != nil {
		log.Printf("[SkinsService] - Failed to parse body: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	updatedSkin, err := config.UpdateSkin(id, skin)
	if err != nil {
		log.Printf("[SkinsService] - Failed to update skin: %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update skin")
	}

	log.Printf("[SkinsService] - Skin updated: %+v", updatedSkin)
	return &updatedSkin, nil
}
