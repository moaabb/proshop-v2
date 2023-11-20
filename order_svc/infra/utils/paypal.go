package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/moaabb/ecommerce/order_svc/domain/order"
	"github.com/moaabb/ecommerce/order_svc/infra/config"
)

func GeneratePaypalToken(cfg *config.Config) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%v/v1/oauth2/token", cfg.PaypalBaseUrl), strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(cfg.PaypalClientId, cfg.PaypalSecretId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var paypalResponseDto order.PayPalAuthDTO
	err = json.NewDecoder(resp.Body).Decode(&paypalResponseDto)
	if err != nil {
		return "", err
	}

	return paypalResponseDto.AccessToken, nil
}
