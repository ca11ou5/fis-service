package service

import (
	"fis/internal/repository"
	"github.com/skbt-ecom/logging"
)

type Service struct {
	repo *repository.Repository
	log  *logging.Logger
}

func NewService(repo *repository.Repository, log *logging.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
