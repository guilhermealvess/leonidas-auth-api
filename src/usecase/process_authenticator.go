package usecase

import (
	"api-auth/src/entity"
	"time"
)

type ProcessAuthenticator struct {
	ProjectRepository entity.ProjectRepository
	AccountRepository entity.AccountRepository
	jwtMaker          JWT
}

type ProcessSignInput struct {
	Credential string
	Key        string
	Email      string
	Password   string
}

type ProcessSignOutput struct {
	Token string
	Error string
}

type ProcessVerifyTokenInput struct {
	Token string
}

type ProcessVerifyTokenOutput struct {
	Eror    string
	Payload Payload
}

func NewProcessAuthenticator(projectRepo entity.ProjectRepository, accountRepo entity.AccountRepository, jwtMaker JWT) *ProcessAuthenticator {
	return &ProcessAuthenticator{
		ProjectRepository: projectRepo,
		AccountRepository: accountRepo,
		jwtMaker:          jwtMaker,
	}
}

func (p *ProcessAuthenticator) Sign(input ProcessSignInput) (*ProcessSignOutput, error) {
	// Validar projeto
	project, err := p.ProjectRepository.FindByCredentials(input.Credential)
	if err != nil {
		return &ProcessSignOutput{}, err
	}

	if project.Key != input.Key {
		return &ProcessSignOutput{}, err
	}

	// Validar conta
	account, err := p.AccountRepository.FindByEmail(input.Email, project.ID.String())
	if err != nil {
		return &ProcessSignOutput{}, err
	}

	// Calcular Hash Password
	if !account.VerifyPassword(input.Password, project.RoudHash, project.HashAlgoritm) {
		return &ProcessSignOutput{}, err
	}

	// Gerar Token
	tokenJWT, err := p.jwtMaker.CreateToken(Payload{
		ID:        account.ID.String(),
		Email:     account.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(60)),
	})

	if err != nil {
		return &ProcessSignOutput{}, err
	}

	// Registrar Login
	account.LastLogin = time.Now().String()
	if p.AccountRepository.Update(*account) != nil {
		return &ProcessSignOutput{}, err
	}

	// Restornar Token
	return &ProcessSignOutput{
		Token: tokenJWT,
	}, nil
}

func (p *ProcessAuthenticator) VerifyToken(input ProcessVerifyTokenInput) (*ProcessVerifyTokenOutput, error) {
	payload, err := p.jwtMaker.Verify(input.Token)

	if err != nil {
		return &ProcessVerifyTokenOutput{
			Eror:    err.Error(),
			Payload: *payload,
		}, err
	}

	return &ProcessVerifyTokenOutput{
		Eror:    "",
		Payload: *payload,
	}, nil
}
