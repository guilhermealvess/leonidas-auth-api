package usecase

import (
	"api-auth/src/entity"
)

type ProjectDtoInput struct {
	Description   string
	HashAlgorithm string
	Name          string
	RoundHash     uint
}

type ProjectDtoOutput struct {
	ApiKey string
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

	project.ApiKey = project.GenerateApiKey()
	project.Secret = project.GenerateSecret()

	_, err := p.Repository.Insert(*project)
	if err == nil {
		output := ProjectDtoOutput{
			ApiKey: project.ApiKey,
		}

		return output, nil
	}

	return ProjectDtoOutput{}, err
}
