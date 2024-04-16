package repository

import (
	"fis/configs"
	"github.com/skbt-ecom/logging"
)

type Repository struct {
	openApi *OpenApiClient
	dadata  *DadataClient
	mongo   *MongoClient
	log     *logging.Logger
}

func NewRepository(cfg *configs.Config, log *logging.Logger) *Repository {
	return &Repository{
		openApi: InitOpenApiClient(cfg),
		dadata:  InitDadataClient(cfg),
		mongo:   InitMongoClient(cfg, log),
		log:     log,
	}
}
