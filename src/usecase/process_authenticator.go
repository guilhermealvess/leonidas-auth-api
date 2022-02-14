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
	Token     string
	ProjectId string
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
	project, err := p.ProjectRepository.FindByCredential(input.Credential)
	if err != nil {
		return &ProcessSignOutput{}, err
	}

	if project.Key != input.Key {
		return &ProcessSignOutput{}, err
	}

	account, err := p.AccountRepository.FindByEmail(input.Email, project.ID)
	if err != nil {
		return &ProcessSignOutput{}, err
	}

	if !account.VerifyPassword(input.Password, project.RoudHash, project.HashAlgoritm) {
		return &ProcessSignOutput{}, err
	}

	tokenJWT, err := p.jwtMaker.CreateToken(Payload{
		ID:        account.ID.String(),
		Email:     account.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(60)),
	}, project.Secret)

	if err != nil {
		return &ProcessSignOutput{}, err
	}

	account.LastLogin = time.Now().String()
	if p.AccountRepository.Update(*account) != nil {
		return &ProcessSignOutput{}, err
	}

	return &ProcessSignOutput{
		Token: tokenJWT,
	}, nil
}

func (p *ProcessAuthenticator) VerifyToken(input ProcessVerifyTokenInput) (*ProcessVerifyTokenOutput, error) {
	project, _ := p.ProjectRepository.FindByID(input.ProjectId)
	payload, err := p.jwtMaker.Verify(input.Token, project.Secret)

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
