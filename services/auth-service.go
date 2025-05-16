package services

import (
	"errors"
	"go-fiber-vercel/config"
	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, "", err
	}

	user.Password = string(hashedPassword)

	newUser, err := config.CreateUser(user)
	if err != nil {
		return models.User{}, "", err
	}

	token, err := helpers.GenerateToken(newUser)
	if err != nil {
		return models.User{}, "", err
	}

	return newUser, token, nil
}
