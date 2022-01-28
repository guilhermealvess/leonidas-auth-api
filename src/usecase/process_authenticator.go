package usecase

import "api-auth/src/entity"

type ProcessAuthenticator struct {
	ProjectRepository entity.ProjectRepository
	AccountRepository entity.AccountRepository
}

func NewProcessAuthenticator(projectRepo entity.ProjectRepository, accountRepo entity.AccountRepository) *ProcessAuthenticator {
	return &ProcessAuthenticator{
		ProjectRepository: projectRepo,
		AccountRepository: accountRepo,
	}
}
