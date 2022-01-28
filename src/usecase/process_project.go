package usecase

import (
	"api-auth/src/entity"
	"time"
)

type ProjectDtoInput struct {
	Description  string
	HashAlgoritm string
	Name         string
	RoudHash     uint
}

type ProjectDtoOutput struct {
	Status int32
	Error  string
	ID     string
}

type ProcessProject struct {
	Repository entity.ProjectRepository
}

func NewProcessProject(repo entity.ProjectRepository) *ProcessProject {
	return &ProcessProject{Repository: repo}
}

func (p *ProcessProject) ExecuteCreateNewProject(projectInput ProjectDtoInput) (ProjectDtoOutput, error) {
	project := entity.NewProject()
	project.Description = projectInput.Description
	project.HashAlgoritm = projectInput.HashAlgoritm
	project.Name = projectInput.Name
	project.RoudHash = projectInput.RoudHash

	err := project.IsValid()
	if err == nil {
		return p.createNewProject(project)
	}

	return ProjectDtoOutput{}, err
}

func (p *ProcessProject) createNewProject(project *entity.Project) (ProjectDtoOutput, error) {

	for {
		credentials := project.GenerateCredential()
		_, err := p.Repository.FindByCredentials(credentials)

		if err != nil {
			project.Crendetials = credentials
			break
		}
	}

	key := project.GenerateKey()
	project.Key = key

	secret := project.GenerateSecret()
	project.Secret = secret

	project.CreatedAt = time.Now()
	project.CreatedBy = "SYSTEM"

	oid, err := p.Repository.Insert(*project)
	if err == nil {
		output := ProjectDtoOutput{
			Error:  "",
			Status: 201,
			ID:     oid.String(),
		}

		return output, nil
	}

	return ProjectDtoOutput{
		Error:  "Não foi possivel criar um projeto",
		Status: 500,
		ID:     "",
	}, err
}
