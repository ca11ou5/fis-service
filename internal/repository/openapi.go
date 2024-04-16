package repository

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fis/configs"
	"fis/internal/entity"
	"fmt"
	"io"
	"net/http"
)

type OpenApiClient struct {
	httpClient *http.Client
	cfg        OpenApiConfig
}

type OpenApiConfig struct {
	BaseUrl    string
	ScoringUrl string
	SecretKey  string
	ClientID   string
}

func InitOpenApiClient(cfg *configs.Config) *OpenApiClient {
	return &OpenApiClient{
		httpClient: http.DefaultClient,
		cfg: OpenApiConfig{
			BaseUrl:    cfg.OpenApiBaseUrl,
			ScoringUrl: cfg.OpenApiScoringUrl,
			SecretKey:  cfg.OpenApiSecretKey,
			ClientID:   cfg.OpenApiClientID,
		},
	}
}

func (c *OpenApiClient) GetScoringStatus(requestId string) (*entity.ScoringStatus, error) {
	path := c.cfg.ScoringUrl + requestId

	method := "GET"
	// TODO: when configuring check for correct url
	url := fmt.Sprintf("%s%s?sourceId=INTERNET_LEADS_API", c.cfg.BaseUrl, path)
	// TODO: when configuring check for correct signature? ecom-api payload before "?", enricher payload all path
	signature := func(payload []byte) string {
		h := hmac.New(sha1.New, []byte(c.cfg.SecretKey))
		h.Write(payload)
		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	}([]byte(path))

	// TODO: NewRequest or NewRequestWithContext
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-id", c.cfg.ClientID)
	req.Header.Add("Signature", signature)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// TODO: answer "why resp have body if method GET"

	var status entity.ScoringStatus
	err = json.Unmarshal(bodyBytes, &status)
	if err != nil {
		return nil, err
	}

	return &status, err
}
