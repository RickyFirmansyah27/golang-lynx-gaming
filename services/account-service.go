package services

import (
	"errors"
	"go-fiber-vercel/helpers"
	"log"
	"net/url"
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

	decodedNickname, err := url.QueryUnescape(nickname)
	if err != nil {
		log.Printf("[AccountService] - Failed to decode nickname: %v", err)
		return "", err
	}

	log.Printf("[AccountService] - Successfully fetched and decoded nickname for gameID: %s, serverID: %s", gameID, serverID)
	return decodedNickname, nil
}
