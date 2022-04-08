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
	coll := m.client.Database(m.database).Collection(collectionName)
	//ctx := getContextWithTimeout(20)
	_, insertResult := coll.InsertOne(context.Background(), document)

	return insertResult
}

func (m *MongoDB) InsertMany() error {
	return nil
}

func (m *MongoDB) FindOne(collectionName string, dataFilter primitive.D) (interface{}, error) {
	var result bson.M
	coll := m.client.Database(m.database).Collection(collectionName)
	err := coll.FindOne(context.Background(), dataFilter).Decode(&result)

	return result, err
}

func (m *MongoDB) FindByID(id string) error {
	return nil
}

func (m *MongoDB) UpdateOne(collectionName string, id primitive.ObjectID, documentFileds primitive.D) error {
	coll := m.client.Database(m.database).Collection(collectionName)
	_, err := coll.UpdateOne(context.Background(), bson.M{"_id": id}, documentFileds)

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
