package usecase

import (
	"api-auth/src/adapter/jwt"
	"api-auth/src/entity"
	"errors"
	"fmt"
)

type AccountDtoInput struct {
	ApiKey                string `json:"apiKey"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	Email                 string `json:"email"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	UrlRedirectActivation string `json:"urlRedirectActivation"`
}

type AccountDtoOutput struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

type ActivationAccountDtoOutput struct {
	Url string `json:"url"`
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
	account.UrlRedirectActivation = input.UrlRedirectActivation
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
		payload := jwt.Payload{
			ID: id,
		}

		tokenJWT, err := p.jwtMaker.CreateToken(payload, account.UID.String())

		if err != nil {
			return &AccountDtoOutput{}, errors.New("Não foi possivel criar uma conta")

		}

		return &AccountDtoOutput{
			ID:   id,
			Link: account.GenerateActivationLink(tokenJWT),
		}, nil
	}

	return &AccountDtoOutput{}, errors.New("Não foi possivel criar uma conta")
}

func (p *ProcessAccount) ActivateAccount(key string) (*ActivationAccountDtoOutput, error) {
	account := entity.NewAccount()

	token, err := account.DecodeActivationKey(key)

	if err != nil {
		return &ActivationAccountDtoOutput{}, err
	}

	payload, err := p.jwtMaker.ParserToken(token)

	if err != nil {
		return &ActivationAccountDtoOutput{}, err
	}

	acc, err := p.Repository.FindByID(payload.ID)

	if err != nil {
		return &ActivationAccountDtoOutput{}, err
	}

	if _, err = p.jwtMaker.Verify(token, acc.UID.String()); err != nil {
		return &ActivationAccountDtoOutput{}, err
	}

	if err = p.Repository.UpdateActived(acc.ID); err != nil {
		return &ActivationAccountDtoOutput{}, err
	}

	return &ActivationAccountDtoOutput{
		Url: acc.UrlRedirectActivation,
	}, nil

}

func (p *ProcessAccount) RefreshActivationLinkAccount(id string) (string, error) {
	account, err := p.Repository.FindByID(id)

	if err != nil {
		return "", nil
	}

	payload := jwt.Payload{ID: id}
	token, err := p.jwtMaker.CreateToken(payload, account.UID.String())

	if err != nil {
		return "", err
	}

	return account.GenerateActivationLink(token), nil
}
