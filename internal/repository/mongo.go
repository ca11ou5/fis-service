package repository

import (
	"context"
	"errors"
	"fis/configs"
	"fis/internal/entity"
	"github.com/skbt-ecom/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var errDocumentNotFound = errors.New("document not found")

type MongoClient struct {
	client *mongo.Client
}

func InitMongoClient(cfg *configs.Config, log *logging.Logger) *MongoClient {
	opts := options.Client().ApplyURI(cfg.MongoUrl)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.WithError(err).Fatalf("failed to connect to mongo on %s", cfg.MongoUrl)
	}

	return &MongoClient{client: client}
}

func (c *MongoClient) GetApplication(requestId string) (*entity.Application, error) {
	appsColl := c.client.Database("app").Collection("apps")
	filter := bson.D{{"apiId", requestId}}

	// TODO: @ecom/schema-app App scheme into go structure
	var application entity.Application
	err := appsColl.FindOne(context.Background(), filter).Decode(&application)
	if err != nil {
		return nil, err
	}

	return &application, nil
}

func (c *MongoClient) UpdateApplicationScoring(appId primitive.ObjectID, status interface{}) error {
	appsColl := c.client.Database("app").Collection("apps")
	filter := bson.D{{"_id", appId}}
	update := bson.M{"$push": bson.M{
		"openApi": []interface{}{
			"scoringGetStatus",
			time.Now().Format(time.RFC3339),
			status,
		},
	}}

	_, err := appsColl.UpdateOne(context.Background(), filter, update)
	if err != nil || errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	return nil
}

func (c *MongoClient) SaveApplication(application *entity.Application) error {
	appsColl := c.client.Database("app").Collection("apps")
	filter := bson.D{{"_id", application.ID}}

	res, err := appsColl.ReplaceOne(context.Background(), filter, application)
	if err != nil || errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	if res.ModifiedCount != 1 {
		return errDocumentNotFound
	}

	return nil
}

func (c *MongoClient) GetUserByUsername(username string) (*entity.UserApplication, error) {
	usersColl := c.client.Database("app").Collection("userapplications")
	filter := bson.D{{"username", username}}

	var userApp entity.UserApplication
	err := usersColl.FindOne(context.Background(), filter).Decode(&userApp)
	if err != nil {
		return nil, err
	}

	return &userApp, nil
}
