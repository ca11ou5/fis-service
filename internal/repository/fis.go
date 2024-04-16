package repository

import (
	"errors"
	"fis/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) GetDadataAddress(query string) *entity.Suggestion {
	sug, err := r.dadata.SendAddressRequest(query)
	if err != nil && !errors.Is(err, errEmptyValue) && !errors.Is(err, errRecordNotFound) {
		r.log.WithError(err).Fatal("failed to send dadata address request")
	}

	return sug
}

func (r *Repository) GetMongoApplication(requestId string) (*entity.Application, error) {
	return r.mongo.GetApplication(requestId)
}

func (r *Repository) UpdateApplicationScoring(appId primitive.ObjectID, status interface{}) error {
	return r.mongo.UpdateApplicationScoring(appId, status)
}

func (r *Repository) GetScoringStatus(requestId string) (*entity.ScoringStatus, error) {
	return r.openApi.GetScoringStatus(requestId)
}

func (r *Repository) SaveApplication(application *entity.Application) error {
	return r.mongo.SaveApplication(application)
}

func (r *Repository) GetUserByUsername(username string) (*entity.UserApplication, error) {
	return r.mongo.GetUserByUsername(username)
}
