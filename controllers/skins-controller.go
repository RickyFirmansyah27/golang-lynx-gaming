package controllers

import (
	"log"

	"go-fiber-vercel/helpers"
	"go-fiber-vercel/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllSkins(c *fiber.Ctx) error {
	log.Printf("[SkinsController] - Incoming request with query params: %v", c.Queries())

	totalData, skins, err := services.GetSkins(c)
	if err != nil {
		log.Printf("[SkinsController] - Failed to fetch skins: %v", err)
		return helpers.Error(c, 400, "Failed to fetch skins", err)
	}

	data := fiber.Map{
		"total_data": totalData,
		"skins":      skins,
	}

	log.Printf("[SkinsController] - Successfully fetched skins")

	return helpers.Success(c, "Successfully fetched skins", data)
}
