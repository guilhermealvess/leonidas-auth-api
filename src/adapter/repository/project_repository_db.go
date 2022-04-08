package repository

import (
	"api-auth/src/entity"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (p *ProjectRepositoryDB) Insert(project entity.Project) (primitive.ObjectID, error) {
	project.ID = primitive.NewObjectID()
	err := p.documentDB.InsertOne(PROJECTS, project)
	if err == nil {
		data, _ := json.Marshal(project)
		p.cache.Set(project.ApiKey, string(data))
	}
	return project.ID, err
}

func (p *ProjectRepositoryDB) FindByCredential(apiKey string) (*entity.Project, error) {
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
