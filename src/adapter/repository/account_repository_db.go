package repository

import (
	"api-auth/src/entity"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ACCOUNT = "accounts"

type AccountRepositoryDB struct {
	entity.AccountRepository
	documentDB DocumentDB
	cache      Cache
}

func NewAccountRepositoryDB(documentDB DocumentDB, cache Cache) *AccountRepositoryDB {
	return &AccountRepositoryDB{
		documentDB: documentDB,
		cache:      cache,
	}
}

func (repo *AccountRepositoryDB) FindByEmail(email string, projectId primitive.ObjectID) (*entity.Account, error) {
	log.Println("FindByEmail, email: %s, projectId: %s", email, projectId)
	account := entity.NewAccount()
	key := email + "-" + projectId.String()
	dataString, err := repo.cache.Get(key)
	if err != nil || dataString == "" {
		data, errorDB := repo.documentDB.FindOne(ACCOUNT, bson.D{{"projectId", projectId}, {"email", email}})

		if errorDB != nil {
			return account, errorDB
		}
		dataByte, _ := json.Marshal(data)
		repo.cache.Set(key, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, account)

		return account, nil
	}
	json.Unmarshal([]byte(dataString), account)

	return account, nil
}

func (repo *AccountRepositoryDB) FindById(id string) (*entity.Account, error) {
	account := entity.NewAccount()
	dataString, err := repo.cache.Get(id)
	if err != nil || dataString == "" {
		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return account, err
		}

		data, errorDB := repo.documentDB.FindOne(ACCOUNT, bson.D{{"_id", oid}})

		if errorDB != nil {
			return account, errorDB
		}

		dataByte, _ := json.Marshal(data)
		repo.cache.Set(id, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, account)

		return account, nil
	}

	json.Unmarshal([]byte(dataString), account)

	return account, nil
}

func (repo *AccountRepositoryDB) Insert(account entity.Account) (primitive.ObjectID, error) {
	account.ID = primitive.NewObjectID()
	err := repo.documentDB.InsertOne(ACCOUNT, account)
	if err == nil {
		data, _ := json.Marshal(account)
		key := account.Email + "-" + account.ProjectId.String()
		repo.cache.Set(key, string(data))
	}
	return account.ID, err
}

func (repo *AccountRepositoryDB) UpdateActived(oid primitive.ObjectID) error {
	return repo.documentDB.UpdateOne(ACCOUNT, oid, bson.D{
		{"$set", bson.D{{"actived", true}, {"updateAt", time.Now()}, {"updateBy", "SYSTEM"}}},
	})
}
