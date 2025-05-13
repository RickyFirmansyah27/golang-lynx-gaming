package controllers

import (
	"log"

	"go-fiber-vercel/helpers"
	"go-fiber-vercel/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllArenas(c *fiber.Ctx) error {
	log.Printf("[ArenasController] - Incoming request with query params: %v", c.Queries())

	totalData, arenas, err := services.GetArenas(c)
	if err != nil {
		log.Printf("[ArenasController] - Failed to fetch arenas: %v", err)
		return helpers.Error(c, 400, "Failed to fetch arenas", err)
	}

	data := fiber.Map{
		"total_data": totalData,
		"arenas":     arenas,
	}

	log.Printf("[ArenasController] - Successfully fetched arenas")

	return helpers.Success(c, "Successfully fetched arenas", data)
}

func UpdateArenas(c *fiber.Ctx) error {
	log.Printf("[ArenasController] - Incoming update request for ID: %s", c.Params("id"))

	updatedArena, err := services.UpdateArenas(c)
	if err != nil {
		log.Printf("[ArenasController] - Failed to update arena: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to update arena", err)
	}

	log.Printf("[ArenasController] - Successfully updated arena: %+v", updatedArena)
	return helpers.Success(c, "Successfully updated arena", updatedArena)
}

func CreateArenas(c *fiber.Ctx) error {
	log.Println("[ArenasController] - Incoming create arena request")

	createdArena, err := services.CreateArenas(c)
	if err != nil {
		log.Printf("[ArenasController] - Failed to create arena: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to create arena", err)
	}

	log.Printf("[ArenasController] - Successfully created arena: %+v", createdArena)
	return helpers.Success(c, "Successfully created arena", createdArena)
}
