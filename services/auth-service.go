package services

import (
	"errors"
	"go-fiber-vercel/config"
	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Login(gameID, serverID, password string) (models.User, string, error) {

	user, err := config.GetUserByGameID(gameID, serverID)
	if err != nil {
		return models.User{}, "", errors.New("invalid game ID or server ID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, "", errors.New("invalid password")
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

func Register(user models.User) (models.User, string, error) {

	validateEmail, err := config.GetUserByEmail(user.Email)
	if err != nil || validateEmail.ID != 0 {
		log.Printf("[AuthService] - Failed to validate email. Email may already exist: %s", user.Email)
		return models.User{}, "", errors.New("email sudah digunakan atau terjadi kesalahan saat validasi")
	}

	validateGameID, err := config.GetUserByGameID(user.GameID, user.ServerID)
	if err != nil || validateGameID.ID != 0 {
		log.Printf("[AuthService] - Failed to validate ID. GameID may already exist: %s", user.GameID)
		return models.User{}, "", errors.New("terjadi kesalahan saat validasi game ID")
	}

	nickname, err := GetAccountDetail(user.GameID, user.ServerID)
	if err != nil {
		log.Printf("[AuthService] - Failed to fetch nickname for GameID: %s, ServerID: %s, error: %v", user.GameID, user.ServerID, err)
		return models.User{}, "", errors.New("failed to fetch nickname from API")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[AuthService] - Failed to hash password: %v", err)
		return models.User{}, "", err
	}

	user.Password = string(hashedPassword)
	user.Nickname = nickname

	newUser, err := config.CreateUser(user)
	if err != nil {
		log.Printf("[AuthService] - Failed to create user: %v", err)
		return models.User{}, "", err
	}

	token, err := helpers.GenerateToken(newUser)
	if err != nil {
		log.Printf("[AuthService] - Failed to generate token: %v", err)
		return models.User{}, "", err
	}

	log.Printf("[AuthService] - Successfully created user: %+v", newUser)
	return newUser, token, nil
}

func CheckAuth(tokenString string) (models.User, error) {
	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		return models.User{}, errors.New("token tidak valid atau sudah kedaluwarsa")
	}

	user := models.User{
		ID:       claims.UserID,
		Name:     claims.Name,
		Email:    claims.Email,
		Nickname: claims.Nickname,
	}

	return user, nil
}
