package entity

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Project struct {
	Name         string    `bson:"name,omitempty"`
	ID           uuid.UUID `bson:"_id,omitempty"`
	Description  string    `bson:"description,omitempty"`
	HashAlgoritm string    `bson:"hashAlgoritm,omitempty"`
	RoudHash     uint      `bson:"roundHash,omitempty"`
	Credential   string    `bson:"credential,omitempty"`
	Key          string    `bson:"key,omitempty"`
	Secret       string    `bson:"secret,omitempty"`
	CreatedBy    string    `bson:"createdBy,omitempty"`
	CreatedAt    time.Time `bson:"createdAt,omitempty"`
	UpdatedBy    string    `bson:"updatedBy,omitempty"`
	UpdatedAt    time.Time `bson:"updatedAt,omitempty"`
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) GenerateKey() string {
	keyLegth := 50
	return p.generateStringRandom(uint(keyLegth))
}

func (p *Project) GenerateCredential() string {
	credentialLegthPart := 32
	firstPartCredential := p.generateStringRandom(uint(credentialLegthPart))
	lastPartCredential := p.generateStringRandom(uint(credentialLegthPart))
	return firstPartCredential + "-" + lastPartCredential
}

func (p *Project) GenerateSecret() string {
	secretLength := 32
	return p.generateStringRandom(uint(secretLength))
}

func (p *Project) generateStringRandom(length uint) string {
	var letterRunes = []rune(alphabet)

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (p *Project) IsValid() error {
	if p.RoudHash < 1 || p.RoudHash > 60 {
		return errors.New("Round Hash over limit permit")
	}

	if p.Description == "" {
		return errors.New("Description invalid")
	}

	p.Description = strings.ToLower(p.Description)

	return nil
}
