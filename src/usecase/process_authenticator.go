package usecase

import (
	"api-auth/src/adapter/jwt"
	"api-auth/src/entity"
	"errors"
	"time"
)

type ProcessAuthenticator struct {
	ProjectRepository entity.ProjectRepository
	AccountRepository entity.AccountRepository
	jwtMaker          jwt.JWT
}

type ProcessSignInput struct {
	ApiKey   string
	Email    string
	Password string
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
	Payload jwt.Payload
}

func NewProcessAuthenticator(projectRepo entity.ProjectRepository, accountRepo entity.AccountRepository, jwtMaker jwt.JWT) *ProcessAuthenticator {
	return &ProcessAuthenticator{
		ProjectRepository: projectRepo,
		AccountRepository: accountRepo,
		jwtMaker:          jwtMaker,
	}
}

func (p *ProcessAuthenticator) Sign(input ProcessSignInput) (*ProcessSignOutput, error) {
	project, err := p.ProjectRepository.FindByApiKey(input.ApiKey)

	if err != nil {
		return &ProcessSignOutput{}, err
	}

	account, err := p.AccountRepository.FindByUsernameAndProject(input.Email, project.ID)
	if err != nil {
		return &ProcessSignOutput{}, err
	}

	if !account.IsActive {
		return &ProcessSignOutput{}, errors.New("Account not active")
	}

	if !account.VerifyPassword(input.Password, project.RoundHash, project.HashAlgoritm) {
		return &ProcessSignOutput{}, errors.New("Password invalid")
	}

	tokenJWT, err := p.jwtMaker.CreateToken(jwt.Payload{
		ID:        account.ID,
		Email:     account.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(60 * time.Minute),
	}, project.Secret)

	if err != nil {
		return &ProcessSignOutput{}, err
	}

	account.LastLogin = time.Now()
	if p.AccountRepository.UpdateLastLogin(account.ID) != nil {
		return &ProcessSignOutput{}, err
	}

	return &ProcessSignOutput{
		Token: tokenJWT,
	}, nil
}

func (p *ProcessAuthenticator) VerifyToken(input ProcessVerifyTokenInput) (*ProcessVerifyTokenOutput, error) {
	project, _ := p.ProjectRepository.FindByID(input.ProjectId)
	payload, err := p.jwtMaker.Verify(input.Token, project.Secret)

	return &ProcessVerifyTokenOutput{
		Payload: *payload,
	}, err
}
