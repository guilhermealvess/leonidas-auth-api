package usecase

import (
	"api-auth/src/entity"
	"time"
)

type ProjectDtoInput struct {
	Description   string
	HashAlgorithm string
	Name          string
	RoundHash     uint
}

type ProjectDtoOutput struct {
	Success bool
	Error   string
	ApiKey  string
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
	project.HashAlgoritm = projectInput.HashAlgorithm
	project.Name = projectInput.Name
	project.RoundHash = projectInput.RoundHash

	err := project.IsValid()
	if err == nil {
		return p.createNewProject(project)
	}

	return ProjectDtoOutput{}, err
}

func (p *ProcessProject) createNewProject(project *entity.Project) (ProjectDtoOutput, error) {

	apiKey := project.GenerateApiKey()
	project.Secret = project.GenerateSecret()
	project.CreatedAt = time.Now()
	project.CreatedBy = "SYSTEM"

	_, err := p.Repository.Insert(*project)
	if err == nil {
		output := ProjectDtoOutput{
			Success: true,
			Error:   "",
			ApiKey:  apiKey,
		}

		return output, nil
	}

	return ProjectDtoOutput{
		Success: false,
		Error:   "Não foi possivel criar um projeto",
		ApiKey:  "",
	}, err
}
