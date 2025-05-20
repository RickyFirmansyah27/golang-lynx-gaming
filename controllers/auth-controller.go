package controllers

import (
	"log"

	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"go-fiber-vercel/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	log.Printf("[AuthController] - Incoming login request")

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
		return helpers.Error(c, fiber.StatusBadRequest, "Failed to register user", err.Error())
	}

	data := fiber.Map{
		"user":  newUser,
		"token": token,
	}

	log.Printf("[AuthController] - Successfully registered user with GameID: %s", newUser.GameID)
	return helpers.Success(c, "Successfully registered user", data)
}

func CheckAuth(c *fiber.Ctx) error {
	log.Printf("[AuthController] - Incoming check auth request")

	var input struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&input); err != nil {
		log.Printf("[AuthController] - Failed to parse body: %v", err)
		return helpers.Error(c, fiber.StatusBadRequest, "Body tidak valid", err)
	}

	if input.Token == "" {
		log.Println("[AuthController] - Token tidak ditemukan di body")
		return helpers.Error(c, fiber.StatusUnauthorized, "Token tidak ditemukan", nil)
	}

	// Verifikasi token
	user, err := services.CheckAuth(input.Token)
	if err != nil {
		log.Printf("[AuthController] - Token validation failed: %v", err)
		return helpers.Success(c, "Token tidak valid atau sudah kedaluwarsa", fiber.Map{
			"nickname": user.Nickname,
		})
	}

	log.Printf("[AuthController] - Successfully logged in user with nickname: %+v", user.Nickname)

	return helpers.Success(c, "Token masih aktif", fiber.Map{
		"nickname": user.Nickname,
	})
}
