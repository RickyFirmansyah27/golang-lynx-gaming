package helpers

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

type CodashopResponse struct {
	ErrorCode          string `json:"errorCode"`
	ConfirmationFields struct {
		Username string `json:"username"`
	} `json:"confirmationFields"`
}

// GetCodashopNickname fetches the nickname from Codashop API based on gameID and serverID
func CheckIDAccount(gameID, serverID string) (string, error) {
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
		log.Printf("[CodashopHelper] - Failed to create request: %v", err)
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[CodashopHelper] - Failed to send request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[CodashopHelper] - API returned non-200 status: %d", resp.StatusCode)
		return "", errors.New("codashop api returned non-200 status")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[CodashopHelper] - Failed to read response body: %v", err)
		return "", err
	}

	var res CodashopResponse
	if err := json.Unmarshal(body, &res); err != nil {
		log.Printf("[CodashopHelper] - Failed to parse JSON: %v", err)
		return "", err
	}

	if res.ErrorCode != "" {
		log.Printf("[CodashopHelper] - API returned error code: %s", res.ErrorCode)
		return "", errors.New("invalid user or zone ID")
	}

	return res.ConfirmationFields.Username, nil
}
