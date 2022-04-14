package db

import (
	"api-auth/src/adapter/repository"
	"context"
	"time"

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

func (m *MongoDB) InsertOne(collectionName string, document interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	coll := m.client.Database(m.database).Collection(collectionName)
	_, insertResult := coll.InsertOne(ctx, document)

	return insertResult
}

func (m *MongoDB) InsertMany() error {
	return nil
}

func (m *MongoDB) FindOne(collectionName string, dataFilter primitive.D) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result bson.M
	coll := m.client.Database(m.database).Collection(collectionName)
	err := coll.FindOne(ctx, dataFilter).Decode(&result)

	return result, err
}

func (m *MongoDB) FindByID(id string) error {
	return nil
}

func (m *MongoDB) UpdateOne(collectionName string, id primitive.ObjectID, documentFileds primitive.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := m.client.Database(m.database).Collection(collectionName)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, documentFileds)

	return err
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
