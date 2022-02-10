package usecase

import "api-auth/src/entity"

type ProcessAuthenticator struct {
	ProjectRepository entity.ProjectRepository
	AccountRepository entity.AccountRepository
	jwtMaker          JWT
}

func NewProcessAuthenticator(projectRepo entity.ProjectRepository, accountRepo entity.AccountRepository, jwtMaker JWT) *ProcessAuthenticator {
	return &ProcessAuthenticator{
		ProjectRepository: projectRepo,
		AccountRepository: accountRepo,
		jwtMaker:          jwtMaker,
	}
}

func (p *ProcessAuthenticator) Sign() {

}
