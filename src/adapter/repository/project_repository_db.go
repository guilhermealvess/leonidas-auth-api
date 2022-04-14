package repository

import (
	"api-auth/src/entity"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectModel struct {
	Name         string             `bson:"name,omitempty"`
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UID          string             `bson:"uid,omitempty"`
	Description  string             `bson:"description,omitempty"`
	HashAlgoritm string             `bson:"hashAlgoritm,omitempty"`
	RoundHash    uint               `bson:"roundHash,omitempty"`
	ApiKey       string             `bson:"apiKey,omitempty"`
	Secret       string             `bson:"secret,omitempty"`
	CreatedBy    string             `bson:"createdBy,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty"`
	UpdatedBy    string             `bson:"updatedBy,omitempty"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty"`
}

const PROJECTS = "projects"

type ProjectRepositoryDB struct {
	entity.ProjectRepository
	documentDB DocumentDB
	cache      Cache
}

func NewProjectRepositoryDB(db DocumentDB, cache Cache) *ProjectRepositoryDB {
	return &ProjectRepositoryDB{
		documentDB: db,
		cache:      cache,
	}
}

func (p *ProjectRepositoryDB) modelToEntity(model ProjectModel) *entity.Project {
	project := entity.NewProject()

	project.Name = model.Name
	project.ID = model.ID.Hex()
	project.ApiKey = model.ApiKey
	project.Description = model.Description
	project.RoundHash = model.RoundHash
	project.Secret = model.Secret
	project.HashAlgoritm = model.HashAlgoritm

	return project
}

func (p *ProjectRepositoryDB) entityToModel(project entity.Project) ProjectModel {
	model := ProjectModel{
		Name:         project.Name,
		UID:          project.UID.String(),
		Description:  project.Description,
		HashAlgoritm: project.HashAlgoritm,
		RoundHash:    project.RoundHash,
		ApiKey:       project.ApiKey,
		Secret:       project.Secret,
	}

	if project.ID != "" {
		model.ID, _ = primitive.ObjectIDFromHex(project.ID)
	}

	return model
}

func (p *ProjectRepositoryDB) Insert(project entity.Project) (string, error) {
	model := p.entityToModel(project)
	model.ID = primitive.NewObjectID()
	model.CreatedAt = time.Now()
	model.CreatedBy = "SYSTEM"

	err := p.documentDB.InsertOne(PROJECTS, model)
	if err == nil {
		data, _ := json.Marshal(project)
		p.cache.Set(project.ApiKey, string(data))
	}
	return project.ID, err
}

func (p *ProjectRepositoryDB) FindByApiKey(apiKey string) (*entity.Project, error) {
	newProject := entity.NewProject()
	dataString, err := p.cache.Get(apiKey)
	if err != nil || dataString == "" {
		data, errorDB := p.documentDB.FindOne(PROJECTS, bson.D{{"apiKey", apiKey}})

		if errorDB != nil {
			return &entity.Project{}, errorDB
		}
		dataByte, _ := json.Marshal(data)
		p.cache.Set(apiKey, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, newProject)

		return newProject, nil
	}

	json.Unmarshal([]byte(dataString), newProject)

	return newProject, nil
}

func (p *ProjectRepositoryDB) FindByID(id string) (*entity.Project, error) {
	newProject := entity.NewProject()
	dataString, err := p.cache.Get(id)
	if err != nil || dataString == "" {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return &entity.Project{}, err
		}
		data, errorDB := p.documentDB.FindOne(PROJECTS, bson.D{{"_id", objectId}})

		if errorDB != nil {
			return &entity.Project{}, errorDB
		}
		dataByte, _ := json.Marshal(data)
		p.cache.Set(id, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, newProject)

		return newProject, nil
	}

	json.Unmarshal([]byte(dataString), newProject)

	return newProject, nil
}
