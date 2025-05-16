package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	"go-fiber-vercel/models"
)

func GetAccountDetail(gameID, serverID string) (string, error) {
	if gameID == "" || serverID == "" {
		log.Printf("[AccountService] - Missing gameID or serverID")
		return "", errors.New("gameID and serverID are required")
	}

	reqBody := models.AccountRequest{
		TypeName: "mobile_legends",
		UserID:   gameID,
		ZoneID:   serverID,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("[AccountService] - Failed to marshal JSON request: %v", err)
		return "", err
	}

	resp, err := http.Post("https://api-cek-id-game-ten.vercel.app/api/check-id-game", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("[AccountService] - Failed to fetch account details: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AccountService] - API returned non-200 status: %d", resp.StatusCode)
		return "", errors.New("failed to fetch account details from API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[AccountService] - Failed to read response body: %v", err)
		return "", err
	}

	var accountDetail models.AccountResponse
	if err := json.Unmarshal(body, &accountDetail); err != nil {
		log.Printf("[AccountService] - Failed to parse JSON: %v", err)
		return "", err
	}

	if !accountDetail.Status {
		log.Printf("[AccountService] - API returned unsuccessful status: %s", accountDetail.Message)
		return "", errors.New(accountDetail.Message)
	}

	decodedNickname, err := url.QueryUnescape(accountDetail.Nickname)
	if err != nil {
		log.Printf("[AccountService] - Failed to decode nickname: %v", err)
		return "", err
	}

	log.Printf("[AccountService] - Successfully fetched nickname for gameID: %s, serverID: %s", gameID, serverID)
	return decodedNickname, nil
}
