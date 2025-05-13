package services

import (
	"log"
	"strconv"

	"go-fiber-vercel/config"
	"go-fiber-vercel/models"

	"github.com/gofiber/fiber/v2"
)

func GetArenas(c *fiber.Ctx) (int, []models.Arenas, error) {
	log.Println("[ArenasService] - Fetching items...", c.Queries())

	queryParams := make(map[string]string)
	c.Context().QueryArgs().VisitAll(func(key, value []byte) {
		queryParams[string(key)] = string(value)
	})

	arenas, totalData, err := config.GetAllArenas(queryParams)
	if err != nil {
		log.Printf("[ArenasService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	log.Printf("[ArenasService] - Successfully fetched %d items", len(arenas))
	return totalData, arenas, nil
}

func UpdateArenas(c *fiber.Ctx) (*models.Arenas, error) {
	log.Println("[ArenasService] - Updating arena...")

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ArenasService] - Invalid ID param: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid ID parameter")
	}

	var arena models.Arenas
	if err := c.BodyParser(&arena); err != nil {
		log.Printf("[ArenasService] - Failed to parse body: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	updatedArena, err := config.UpdateArena(id, arena)
	if err != nil {
		log.Printf("[ArenasService] - Failed to update arena: %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update arena")
	}

	log.Printf("[ArenasService] - Arena updated: %+v", updatedArena)
	return &updatedArena, nil
}

func CreateArenas(c *fiber.Ctx) (*models.Arenas, error) {
	log.Println("[ArenasService] - Creating new arena...")

	var arena models.Arenas
	if err := c.BodyParser(&arena); err != nil {
		log.Printf("[ArenasService] - Failed to parse body: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	createdArena, err := config.CreateArena(arena)
	if err != nil {
		log.Printf("[ArenasService] - Failed to create arena: %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create arena")
	}

	log.Printf("[ArenasService] - Arena created: %+v", createdArena)
	return &createdArena, nil
}
