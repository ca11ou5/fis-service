package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fis/configs"
	"fis/internal/entity"
	"io"
	"net/http"
)

var errEmptyValue = errors.New("value is empty")
var errRecordNotFound = errors.New("record not found")

type DadataClient struct {
	httpClient *http.Client
	cfg        DadataConfig
}

type DadataConfig struct {
	BaseUrl    string
	AddressUrl string
}

func InitDadataClient(cfg *configs.Config) *DadataClient {
	return &DadataClient{
		httpClient: http.DefaultClient,
		cfg: DadataConfig{
			BaseUrl:    cfg.DadataBaseUrl,
			AddressUrl: cfg.DadataAddressUrl,
		},
	}
}

func (c *DadataClient) SendAddressRequest(query string) (*entity.Suggestion, error) {
	emptySug := entity.Suggestion{}
	if query == "" {
		return &emptySug, errEmptyValue
	}

	dadataReq := entity.DadataRequest{
		Query: query,
		Count: 1,
	}

	bodyBytes, err := json.Marshal(dadataReq)
	if err != nil {
		return &emptySug, err
	}

	req, err := http.NewRequest("POST", c.cfg.BaseUrl+c.cfg.AddressUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return &emptySug, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &emptySug, err
	}
	defer resp.Body.Close()

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return &emptySug, err
	}

	var dadataResp entity.DadataAddressResponse
	err = json.Unmarshal(bodyBytes, &dadataResp)
	if err != nil {
		return &emptySug, err
	}

	if len(dadataResp.Suggestions) != 1 {
		return &emptySug, errRecordNotFound
	}

	return &dadataResp.Suggestions[0], nil
}
