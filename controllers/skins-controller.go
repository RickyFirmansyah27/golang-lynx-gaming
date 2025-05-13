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

func UpdateSkins(c *fiber.Ctx) error {
	log.Printf("[SkinsController] - Incoming update request for ID: %s", c.Params("id"))

	updatedSkin, err := services.UpdateSkins(c)
	if err != nil {
		log.Printf("[SkinsController] - Failed to update skin: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to update skin", err)
	}

	log.Printf("[SkinsController] - Successfully updated skin: %+v", updatedSkin)
	return helpers.Success(c, "Successfully updated skin", updatedSkin)
}

func CreateSkins(c *fiber.Ctx) error {
	log.Println("[SkinsController] - Incoming create skin request")

	createdSkin, err := services.CreateSkins(c)
	if err != nil {
		log.Printf("[SkinsController] - Failed to create skin: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to create skin", err)
	}

	log.Printf("[SkinsController] - Successfully created skin: %+v", createdSkin)
	return helpers.Success(c, "Successfully created skin", createdSkin)
}
