package repository

import (
	"api-auth/src/entity"
	"encoding/json"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (p *ProjectRepositoryDB) Insert(project entity.Project) (uuid.UUID, error) {
	oid, err := p.documentDB.InsertOne(PROJECTS, project)
	if err == nil {
		project.ID = oid
		data, _ := json.Marshal(project)
		p.cache.Set(project.Credential, string(data))
	}
	return oid, err
}

func (p *ProjectRepositoryDB) FindByCredential(credential string) (*entity.Project, error) {
	newProject := entity.NewProject()
	dataString, err := p.cache.Get(credential)
	if err != nil || dataString == "" {
		data, errorDB := p.documentDB.FindOne(PROJECTS, bson.D{{"credential", credential}})

		if errorDB != nil {
			return &entity.Project{}, errorDB
		}
		dataByte, _ := json.Marshal(data)
		p.cache.Set(credential, string(dataByte))

		j, _ := json.Marshal(data)
		json.Unmarshal(j, newProject)

		return newProject, nil
	}

	json.Unmarshal([]byte(dataString), newProject)

	return newProject, nil
}
