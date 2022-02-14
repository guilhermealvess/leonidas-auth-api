package db

import (
	"api-auth/src/adapter/repository"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	repository.DocumentDB
	client   mongo.Client
	database string
}

func NewMongoDBInstance(client mongo.Client, database string) *MongoDB {
	return &MongoDB{
		client:   client,
		database: database,
	}
}

func (m *MongoDB) InsertOne(collectionName string, document interface{}) (uuid.UUID, error) {
	coll := m.client.Database(m.database).Collection(collectionName)
	//ctx := getContextWithTimeout(20)
	result, insertResult := coll.InsertOne(context.TODO(), document)

	oid := result.InsertedID.(primitive.ObjectID)

	return uuid.MustParse(oid.String()), insertResult
}

func (m *MongoDB) InsertMany() error {
	return nil
}

func (m *MongoDB) FindOne(collectionName string, dataFilter primitive.D) (interface{}, error) {
	var result bson.M
	coll := m.client.Database(m.database).Collection(collectionName)
	err := coll.FindOne(context.TODO(), dataFilter).Decode(&result)

	return result, err
}

func (m *MongoDB) FindByID(id string) error {
	return nil
}

func (m *MongoDB) UpdateOne() error {
	return nil
}

func (m *MongoDB) UpdateMany() error {
	return nil
}

func (m *MongoDB) Delete(id string) error {
	return nil
}

func getContextWithTimeout(duration time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return ctx
}
