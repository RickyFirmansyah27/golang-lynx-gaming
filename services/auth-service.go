package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"go-fiber-vercel/config"
	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string) (string, error) {
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", err
	}
	salt := hex.EncodeToString(saltBytes)
	hash := pbkdf2.Key([]byte(password), saltBytes, 1000, 64, sha512.New)
	return salt + ":" + hex.EncodeToString(hash), nil
}

// VerifyPassword membandingkan password raw dengan password hashed
func VerifyPassword(passwordRaw, hashed string) bool {
	parts := strings.Split(hashed, ":")
	if len(parts) != 2 {
		return false
	}
	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false
	}
	storedHash := parts[1]
	hash := pbkdf2.Key([]byte(passwordRaw), salt, 1000, 64, sha512.New)
	return hex.EncodeToString(hash) == storedHash
}

func RegisterUser(name, email, password string) (*models.User, error) {
	existingUser, err := config.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser.ID != 0 {
		return nil, errors.New("email is already taken")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		GameID:   "default",
		ServerID: "default",
	}

	createdUser, err := config.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func LoginUser(email, password string) (string, *models.User, error) {
	user, err := config.GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("user not found")
	}
	if !VerifyPassword(password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}
