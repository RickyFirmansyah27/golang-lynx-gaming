package services

import (
	"errors"
	"log"

	"go-fiber-vercel/helpers"
)

func GetAccountDetail(gameID, serverID string) (string, error) {
	if gameID == "" || serverID == "" {
		log.Printf("[AccountService] - Missing gameID or serverID")
		return "", errors.New("gameID and serverID are required")
	}

	nickname, err := helpers.CheckIDAccount(gameID, serverID)
	if err != nil {
		log.Printf("[AccountService] - Failed to get nickname from Codashop: %v", err)
		return "", err
	}

	log.Printf("[AccountService] - Successfully fetched nickname for gameID: %s, serverID: %s", gameID, serverID)
	return nickname, nil
}
