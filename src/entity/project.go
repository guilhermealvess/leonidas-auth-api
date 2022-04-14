package entity

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

const (
	alphabet            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ROUND_LIMIT         = 60
	secretLength        = 32
	credentialLegthPart = 32
	keyLegth            = 50
)

type Project struct {
	Name         string    `json:"name"`
	UID          uuid.UUID `json:"uid"`
	ID           string    `json:"_id"`
	Description  string    `json:"description"`
	HashAlgoritm string    `json:"hashAlgoritm"`
	RoundHash    uint      `json:"roundHash"`
	ApiKey       string    `json:"apiKey"`
	Secret       string    `json:"secret"`
}

func NewProject() *Project {
	return &Project{
		UID: uuid.New(),
	}
}

func (p *Project) GenerateApiKey() string {
	return uuid.NewString()
}

func (p *Project) GenerateSecret() string {
	return uuid.NewString()
}

func (p *Project) IsValid() error {
	if p.RoundHash < 1 || p.RoundHash > uint(ROUND_LIMIT) {
		return errors.New("Round Hash over limit permit")
	}

	if p.Description == "" {
		return errors.New("Description invalid")
	}

	p.Description = strings.ToLower(p.Description)

	return nil
}
