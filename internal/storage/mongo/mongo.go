package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

type Document struct {
	Url   string `bson:"url"`
	Alias string `bson:"alias"`
}

func New(ctx context.Context, connString string) (*Storage, error) {
	const op = "storage.mongo.New"
	client, err := mongo.Connect(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	collection := client.Database("url-shortener").Collection("url")
	return &Storage{
		client:     client,
		collection: collection,
	}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) error {
	const op = "storage.mongo.SaveURL"
	doc := Document{
		Url:   urlToSave,
		Alias: alias,
	}
	_, err := s.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.mongo.GetURL"
	var doc Document
	err := s.collection.FindOne(context.TODO(), bson.M{"alias": alias}).Decode(&doc)
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}
	return doc.Url, nil
}

func (s *Storage) DeleteUrl(alias string) error {
	const op = "storage.mongo.DeleteURL"
	_, err := s.collection.DeleteOne(context.TODO(), bson.M{"alias": alias})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) Close() error {
	return s.client.Disconnect(context.TODO())
}
