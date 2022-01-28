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

	oid, err := p.documentDB.InsertOne(PROJECTS, project)
	if err == nil {
		project.ID = oid
		data, _ := json.Marshal(project)
		p.cache.Set(project.Crendetials, string(data))
	}
	return oid, err
}

func (p *ProjectRepositoryDB) FindByCredentials(credentials string) (*entity.Project, error) {
	newProject := entity.NewProject()
	dataString, err := p.cache.Get(credentials)
	if err != nil || dataString == "" {
		data, errorDB := p.documentDB.FindOne(PROJECTS, bson.D{{"credentials", credentials}})

		if errorDB != nil {
			return &entity.Project{}, errorDB
		}
		dataByte, _ := json.Marshal(data)
		p.cache.Set(credentials, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, newProject)

		return newProject, nil
	}

	json.Unmarshal([]byte(dataString), newProject)

	return newProject, nil
}
