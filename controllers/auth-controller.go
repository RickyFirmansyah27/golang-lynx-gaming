package controllers

import (
	"log"

	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"go-fiber-vercel/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	log.Printf("[AuthController] - Incoming login request with body: %s", c.Body())

	var input struct {
		GameID   string `json:"gameId"`
		ServerID string `json:"serverId"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		log.Printf("[AuthController] - Failed to parse login request body: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	user, token, err := services.Login(input.GameID, input.ServerID, input.Password)
	if err != nil {
		log.Printf("[AuthController] - Failed to login: %v", err)
		return helpers.Error(c, fiber.StatusUnauthorized, "Login failed", err)
	}

	data := fiber.Map{
		"user":  user,
		"token": token,
	}

	log.Printf("[AuthController] - Successfully logged in user with GameID: %s", input.GameID)
	return helpers.Success(c, "Successfully logged in", data)
}

func Register(c *fiber.Ctx) error {
	log.Printf("[AuthController] - Incoming register request with body: %s", c.Body())

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Printf("[AuthController] - Failed to parse register request body: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	newUser, token, err := services.Register(user)
	if err != nil {
		log.Printf("[AuthController] - Failed to register user: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to register user", err)
	}

	data := fiber.Map{
		"user":  newUser,
		"token": token,
	}

	log.Printf("[AuthController] - Successfully registered user with GameID: %s", newUser.GameID)
	return helpers.Success(c, "Successfully registered user", data)
}
