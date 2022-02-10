package repository

import (
	"api-auth/src/entity"
	"encoding/json"

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

func (repo *AccountRepositoryDB) FindByEmail(email string, projectId string) (*entity.Account, error) {
	account := entity.NewAccount()
	key := email + "-" + projectId
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

func (repo *AccountRepositoryDB) Insert(account entity.Account) (primitive.ObjectID, error) {
	oid, err := repo.documentDB.InsertOne(ACCOUNT, account)
	if err == nil {
		account.ID = oid

		data, _ := json.Marshal(account)
		key := account.Email + "-" + account.ProjectId.String()
		repo.cache.Set(key, string(data))
	}
	return oid, err
}
