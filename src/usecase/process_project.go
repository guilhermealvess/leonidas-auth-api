package usecase

import (
	"api-auth/src/entity"
	"time"
)

type ProjectDtoInput struct {
	Description  string
	HashAlgoritm string
	Name         string
	RoundHash    uint
}

type ProjectDtoOutput struct {
	Status     int32
	Error      string
	ID         string
	Credential string
	Key        string
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
	project.RoundHash = projectInput.RoundHash

	err := project.IsValid()
	if err == nil {
		return p.createNewProject(project)
	}

	return ProjectDtoOutput{}, err
}

func (p *ProcessProject) createNewProject(project *entity.Project) (ProjectDtoOutput, error) {

	for {
		credential := project.GenerateCredential()
		_, err := p.Repository.FindByCredential(credential)

		if err != nil {
			project.Credential = credential
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
			Error:      "",
			Status:     201,
			ID:         oid.Hex(),
			Credential: project.Credential,
			Key:        project.Key,
		}

		return output, nil
	}

	return ProjectDtoOutput{
		Error:      "Não foi possivel criar um projeto",
		Status:     500,
		ID:         "",
		Credential: "",
		Key:        "",
	}, err
}
