package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cache interface {
	Get(key string) (string, error)

	Set(key string, value string) error

	SetExpirationSecound(key string, value string, duration time.Duration) error

	Delete(key string)
}

type DocumentDB interface {
	InsertOne(collectionName string, document interface{}) error

	InsertMany() error

	FindOne(collectionName string, dataFilter primitive.D) (interface{}, error)

	FindByID(id string) error

	UpdateOne() error

	UpdateMany() error

	Delete(id string) error
}
