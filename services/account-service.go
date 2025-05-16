package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetAccountDetail(gameID, serverID string) (string, error) {
	if gameID == "" || serverID == "" {
		log.Printf("[AccountService] - Missing gameID or serverID")
		return "", errors.New("gameID and serverID are required")
	}

	form := url.Values{}
	form.Add("voucherPricePoint.id", "5199")
	form.Add("voucherPricePoint.price", "68543.0000")
	form.Add("voucherPricePoint.variablePrice", "0")
	form.Add("user.userId", gameID)
	form.Add("user.zoneId", serverID)
	form.Add("voucherTypeName", "MOBILE_LEGENDS")
	form.Add("shopLang", "id_ID")

	headers := map[string]string{
		"Host":            "order-sg.codashop.com",
		"Accept-Language": "id-ID",
		"Origin":          "https://www.codashop.com",
		"Referer":         "https://www.codashop.com/",
		"Content-Type":    "application/x-www-form-urlencoded",
	}

	req, err := http.NewRequest("POST", "https://order-sg.codashop.com/initPayment.action", strings.NewReader(form.Encode()))
	if err != nil {
		log.Printf("[AccountService] - Failed to create request: %v", err)
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[AccountService] - Failed to send request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AccountService] - API returned non-200 status: %d", resp.StatusCode)
		return "", errors.New("failed to fetch account details from Codashop API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[AccountService] - Failed to read response body: %v", err)
		return "", err
	}

	var responseData struct {
		ErrorCode          string `json:"errorCode"`
		ConfirmationFields struct {
			Username string `json:"username"`
		} `json:"confirmationFields"`
	}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("[AccountService] - Failed to parse JSON response: %v", err)
		return "", err
	}

	if responseData.ErrorCode != "" {
		log.Printf("[AccountService] - API returned error code: %s", responseData.ErrorCode)
		return "", errors.New("invalid user or zone ID")
	}

	nickname := responseData.ConfirmationFields.Username
	decodedNickname, err := url.QueryUnescape(nickname)
	if err != nil {
		log.Printf("[AccountService] - Failed to decode nickname: %v", err)
		return "", err
	}

	log.Printf("[AccountService] - Successfully fetched nickname for gameID: %s, serverID: %s", gameID, serverID)
	return decodedNickname, nil
}
