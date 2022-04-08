package usecase

import (
	"api-auth/src/adapter/jwt"
	"api-auth/src/entity"
	"errors"

	"github.com/google/uuid"
)

type AccountDtoInput struct {
	ApiKey                string
	Name                  string
	LastName              string
	Email                 string
	Password              string
	UrlRedirectActivation string
}

type AccountDtoOutput struct {
	Success bool
	Error   string
	ID      string
	Link    string
}

type ActivationAccountDtoOutput struct {
	Success bool
	Error   string
	Url     string
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
	account := entity.NewAccount()

	project, err := p.projectRepository.FindByApiKey(input.ApiKey)
	
	if err != nil {
		return &AccountDtoOutput{}, errors.New("Credential invalid")
	}

	account.Name = input.Name
	account.LastName = input.LastName
	account.Email = input.Email
	account.Password = input.Password
	account.ProjectId = project.ID
	account.UrlRedirectLaterActivation = input.UrlRedirectActivation
	account.Activated = false

	err = account.IsValid()
	if err != nil {
		return &AccountDtoOutput{
			Success: false,
			Error:   err.Error(),
			ID:      "",
			Link:    "",
		}, err
	}

	err = account.SavePassword(input.Password, project.HashAlgoritm, project.RoundHash)
	if err != nil {
		return &AccountDtoOutput{
			Success: false,
			Error:   err.Error(),
			ID:      "",
			Link:    "",
		}, err
	}

	return p.createAccount(*account)
}

func (p *ProcessAccount) createAccount(account entity.Account) (*AccountDtoOutput, error) {
	_, err := p.Repository.FindByEmail(account.Email, account.ProjectId)
	if err == nil {
		return &AccountDtoOutput{}, errors.New("Account not enable")
	}

	activationSecret := uuid.NewString()
	account.ActivationSecret = activationSecret
	oid, err := p.Repository.Insert(account)

	if err == nil {
		payload := jwt.Payload{
			ID: oid.Hex(),
		}

		tokenJWT, err := p.jwtMaker.CreateToken(payload, activationSecret)

		if err != nil {
			return &AccountDtoOutput{
				Error:   "Não foi possivel criar uma conta",
				Success: false,
				ID:      "",
				Link:    "",
			}, err

		}

		output := &AccountDtoOutput{
			Error:   "",
			Success: true,
			ID:      oid.Hex(),
			Link:    account.GenerateActivationLink(tokenJWT),
		}

		return output, nil
	}

	return &AccountDtoOutput{
		Error:   "Não foi possivel criar uma conta",
		Success: false,
		ID:      "",
		Link:    "",
	}, err
}

func (p *ProcessAccount) ActivateAccount(key string) *ActivationAccountDtoOutput {
	account := entity.NewAccount()

	token, err := account.DecodeActivationKey(key)

	if err != nil {
		return activateAccountFail(err)
	}

	payload, err := p.jwtMaker.ParserToken(token)

	if err != nil {
		return activateAccountFail(err)
	}

	acc, err := p.Repository.FindByID(payload.ID)

	if err != nil {
		return activateAccountFail(err)
	}

	if _, err = p.jwtMaker.Verify(token, acc.ActivationSecret); err != nil {
		return activateAccountFail(err)
	}

	if err = p.Repository.UpdateActived(acc.ID); err != nil {
		return activateAccountFail(err)
	}

	return activateAccountFail(err)

}

func activateAccountFail(err error) *ActivationAccountDtoOutput {
	return &ActivationAccountDtoOutput{
		Success: false,
		Error:   err.Error(),
		Url:     "",
	}
}

func (p *ProcessAccount) RefreshActivationLinkAccount(id string) (string, error) {
	account, err := p.Repository.FindByID(id)

	if err != nil {
		return "", nil
	}

	payload := jwt.Payload{ID: id}
	token, err := p.jwtMaker.CreateToken(payload, account.ActivationSecret)

	if err != nil {
		return "", err
	}

	return account.GenerateActivationLink(token), nil
}
