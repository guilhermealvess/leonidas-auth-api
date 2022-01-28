package entity

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Project struct {
	Name         string             `json: name`
	ID           primitive.ObjectID `json: _id`
	Description  string             `json: description`
	HashAlgoritm string             `json: hashAlgoritm`
	RoudHash     uint               `json: roundHash`
	Crendetials  string             `json: credentials`
	Key          string             `json: key`
	Secret       string             `json: secret`
	CreatedBy    string             `json: createdBy`
	CreatedAt    time.Time          `json: createdAt`
	UpdatedBy    string             `json: updatedBy`
	UpdatedAt    time.Time          `json: updatedAt`
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) GenerateKey() string {
	keyLegth := 50
	return p.generateStringRandom(uint(keyLegth))
}

func (p *Project) GenerateCredential() string {
	credentialsLegthPart := 32
	firstPartCredentials := p.generateStringRandom(uint(credentialsLegthPart))
	lastPartCredentials := p.generateStringRandom(uint(credentialsLegthPart))
	return firstPartCredentials + "-" + lastPartCredentials
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
