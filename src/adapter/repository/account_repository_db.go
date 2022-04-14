package repository

import (
	"api-auth/src/entity"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ACCOUNT    = "accounts"
	ID         = "_id"
	PROJECT_ID = "projectId"
	USERNAME   = "username"
	ACTIVED    = "activated"
	LAST_LOGIN = "lastLogin"
	PASSWORD   = "password"
	CREATED_AT = "createdAt"
	CREATED_BY = "createdBy"
	UPDATED_AT = "updatedAt"
	UPDATED_BY = "updatedBy"
	SYSTEM     = "SYSTEM"
)

type AccountRepositoryDB struct {
	entity.AccountRepository

	documentDB DocumentDB
	cache      Cache
}

type AccountModel struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ProjectID     primitive.ObjectID `bson:"projectId,omitempty"`
	UID           string             `bson:"uid,omitempty"`
	FirstName     string             `bson:"firstName,omitempty"`
	LastName      string             `bson:"lastName,omitempty"`
	Email         string             `bson:"email,omitempty"`
	Username      string             `bson:"username,omitempty"`
	Password      string             `bson:"password,omitempty"`
	LastLogin     time.Time          `bson:"lastLogin,omitempty"`
	IsActive      bool               `bson:"activated"`
	VerifiedEmail bool               `bson:"verifiedEmail,omitempty"`
	ActivedAt     time.Time          `bson:"activedAt,omitempty"`
	createdAt     time.Time          `bson:"createdAt,omitempty"`
	createdBy     string             `bson:"createdBy,omitempty"`
	updatedAt     time.Time          `bson:"updatedAt,omitempty"`
	updatedBy     string             `bson:"updatedBy,omitempty"`
}

func NewAccountModel() *AccountModel {
	return &AccountModel{
		ID: primitive.NewObjectID(),
	}
}

func NewAccountRepositoryDB(documentDB DocumentDB, cache Cache) *AccountRepositoryDB {
	return &AccountRepositoryDB{
		documentDB: documentDB,
		cache:      cache,
	}
}

func (repo *AccountRepositoryDB) modelToEntity(model AccountModel) *entity.Account {
	account := &entity.Account{
		ID:            model.ID.Hex(),
		ProjectID:     model.ProjectID.Hex(),
		UID:           uuid.MustParse(model.UID),
		FirstName:     model.FirstName,
		LastName:      model.LastName,
		Email:         model.Email,
		Username:      model.Username,
		Password:      model.Password,
		VerifiedEmail: model.VerifiedEmail,
		IsActive:      model.IsActive,
		ActivedAt:     model.ActivedAt,
		LastLogin:     model.LastLogin,
	}

	return account
}

func (p *AccountRepositoryDB) entityToModel(account entity.Account) *AccountModel {
	projectID, _ := primitive.ObjectIDFromHex(account.ProjectID)

	accountID := primitive.NewObjectID()
	if account.ID != "" {
		accountID, _ = primitive.ObjectIDFromHex(account.ID)
	}

	model := &AccountModel{
		ID:            accountID,
		ProjectID:     projectID,
		UID:           account.UID.String(),
		FirstName:     account.FirstName,
		LastName:      account.LastName,
		Email:         account.Email,
		Username:      account.Username,
		Password:      account.Password,
		LastLogin:     account.LastLogin,
		IsActive:      account.IsActive,
		VerifiedEmail: account.VerifiedEmail,
		ActivedAt:     account.ActivedAt,
	}

	return model
}

func (repo *AccountRepositoryDB) FindByUsernameAndProject(username string, projectID string) (*entity.Account, error) {
	log.Println("FindByUsernameAndProject, username: %s, projectID: %s", username, projectID)
	var account entity.Account
	var model AccountModel

	key := projectID + ":" + username
	dataString, err := repo.cache.Get(key)

	if err != nil || dataString == "" {
		oid, _ := primitive.ObjectIDFromHex(projectID)
		data, errorDB := repo.documentDB.FindOne(ACCOUNT, bson.D{{PROJECT_ID, oid}, {USERNAME, username}})

		if errorDB != nil {
			return &account, errorDB
		}

		dataByte, _ := json.Marshal(data)
		repo.cache.Set(key, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, &model)

		account = *repo.modelToEntity(model)
		return &account, nil
	}
	json.Unmarshal([]byte(dataString), &model)

	return repo.modelToEntity(model), nil
}

func (repo *AccountRepositoryDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var model AccountModel

	dataString, err := repo.cache.Get(id)

	if err != nil || dataString == "" {
		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return &account, err
		}

		data, errorDB := repo.documentDB.FindOne(ACCOUNT, bson.D{{ID, oid}})

		if errorDB != nil {
			return &account, errorDB
		}

		dataByte, _ := json.Marshal(data)
		repo.cache.Set(id, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, &model)

		return repo.modelToEntity(model), nil
	}

	json.Unmarshal([]byte(dataString), &model)

	return repo.modelToEntity(model), nil
}

func (repo *AccountRepositoryDB) Insert(account entity.Account) (string, error) {
	acc := repo.entityToModel(account)
	acc.createdAt = time.Now()
	acc.createdBy = SYSTEM

	err := repo.documentDB.InsertOne(ACCOUNT, acc)

	if err == nil {
		data, _ := json.Marshal(account)
		key := account.UID.String()
		repo.cache.Set(key, string(data))
	}
	return acc.ID.Hex(), err
}

func (repo *AccountRepositoryDB) UpdateActived(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	return repo.documentDB.UpdateOne(ACCOUNT, oid, bson.D{
		{"$set", bson.D{{ACTIVED, true}, {UPDATED_AT, time.Now()}, {UPDATED_BY, SYSTEM}}},
	})
}

func (repo *AccountRepositoryDB) UpdateLastLogin(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	now := time.Now()
	return repo.documentDB.UpdateOne(ACCOUNT, oid, bson.D{
		{"$set", bson.D{{LAST_LOGIN, now}, {UPDATED_AT, now}, {UPDATED_BY, SYSTEM}}},
	})
}

func (repo *AccountRepositoryDB) UpdatePassword(id string, password string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	now := time.Now()
	return repo.documentDB.UpdateOne(ACCOUNT, oid, bson.D{
		{"$set", bson.D{{PASSWORD, password}, {UPDATED_AT, now}, {UPDATED_BY, SYSTEM}}},
	})
}
