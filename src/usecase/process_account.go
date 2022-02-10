package usecase

import (
	"api-auth/src/entity"
	"errors"
)

type AccountDtoInput struct {
	Credentials string
	Key         string
	Name        string
	LastName    string
	Email       string
	Password    string
}

type AccountDtoOutput struct {
	Status int32
	Error  string
	ID     string
}

type ProcessAccount struct {
	Repository        entity.AccountRepository
	projectRepository entity.ProjectRepository
}

func NewProcessAccount(repository entity.AccountRepository, projectRepository entity.ProjectRepository) *ProcessAccount {
	return &ProcessAccount{
		Repository:        repository,
		projectRepository: projectRepository,
	}
}

func (p *ProcessAccount) ExecuteCreateNewAccount(input AccountDtoInput) (*AccountDtoOutput, error) {
	account := entity.NewAccount()

	project, err := p.projectRepository.FindByCredentials(input.Credentials)
	if err != nil {
		return &AccountDtoOutput{}, errors.New("Credentials invalid")
	}

	account.Name = input.Name
	account.LastName = input.LastName
	account.Email = input.Email
	account.Password = input.Password
	account.ProjectId = project.ID

	err = account.IsValid()
	if err != nil {
		return &AccountDtoOutput{}, err
	}

	return p.createAccount(*account)
}

func (p *ProcessAccount) createAccount(account entity.Account) (*AccountDtoOutput, error) {
	_, err := p.Repository.FindByEmail(account.Email, account.ProjectId.String())
	if err == nil {
		return &AccountDtoOutput{}, errors.New("Account not enable")
	}

	oid, err := p.Repository.Insert(account)
	if err == nil {
		output := &AccountDtoOutput{
			Error:  "",
			Status: 201,
			ID:     oid.String(),
		}

		return output, nil
	}

	return &AccountDtoOutput{
		Error:  "Não foi possivel criar uma conta",
		Status: 500,
		ID:     "",
	}, err
}
