package usecase

import (
	"api-auth/src/adapter/jwt"
	"api-auth/src/entity"
	"errors"
	"fmt"
)

type AccountDtoInput struct {
	ApiKey    string `json:"apiKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type AccountDtoOutput struct {
	ID string `json:"id"`
}

type ActivationAccountDtoInput struct {
	ApiKey   string `json:"apiKey"`
	Username string `json:"username"`
}

type SetNewPassowrdDtoInput struct {
	ApiKey      string `json:"apiKey"`
	Username    string `json:"username"`
	Token       string `json:"token"`
	NewPassword string `json:"newPassword"`
}

type ProcessAccount struct {
	Repository        entity.AccountRepository
	projectRepository entity.ProjectRepository
	jwtMaker          jwt.JWT
}

func NewProcessAccount(repository entity.AccountRepository, projectRepository entity.ProjectRepository) *ProcessAccount {
	return &ProcessAccount{
		Repository:        repository,
		projectRepository: projectRepository,
		jwtMaker:          jwt.NewJWTMaker(),
	}
}

func (p *ProcessAccount) ExecuteCreateNewAccount(input AccountDtoInput) (*AccountDtoOutput, error) {
	project, err := p.projectRepository.FindByApiKey(input.ApiKey)

	if err != nil {
		return &AccountDtoOutput{}, errors.New("Credential invalid")
	}

	account := entity.NewAccount()
	account.FirstName = input.FirstName
	account.LastName = input.LastName
	account.Email = input.Email
	account.Username = input.Username
	account.Password = input.Password
	account.ProjectID = project.ID
	account.IsActive = false

	err = account.IsValid()
	if err != nil {
		return &AccountDtoOutput{}, err
	}

	a, err := p.Repository.FindByUsernameAndProject(account.Username, account.ProjectID)
	if err == nil {
		fmt.Print(a)
		return &AccountDtoOutput{}, errors.New("Username not available")
	}

	err = account.SavePassword(input.Password, project.HashAlgoritm, project.RoundHash)
	if err != nil {
		return &AccountDtoOutput{}, err
	}

	id, err := p.Repository.Insert(*account)

	if err == nil {
		return &AccountDtoOutput{
			ID:    id,
		}, nil
	}

	return &AccountDtoOutput{}, errors.New("Não foi possivel criar uma conta")
}

func (p *ProcessAccount) ActivateAccount(input ActivationAccountDtoInput) error {
	project, err := p.projectRepository.FindByApiKey(input.ApiKey)

	if err != nil {
		return errors.New("Credential invalid")
	}

	account, err := p.Repository.FindByUsernameAndProject(input.Username, project.ID)

	if err != nil {
		return err
	}

	account.ActivedAccount()

	return p.Repository.UpdateActived(account.ID)
}

func (p *ProcessAccount) SetNewPassword(input SetNewPassowrdDtoInput) error {
	project, err := p.projectRepository.FindByApiKey(input.ApiKey)

	if err != nil {
		return errors.New("Credential invalid")
	}

	account, err := p.Repository.FindByUsernameAndProject(input.Username, project.ID)

	if err != nil {
		return err
	}

	account.SavePassword(input.NewPassword, project.HashAlgoritm, project.RoundHash)

	return p.Repository.UpdatePassword(account.ID, account.Password)
}
