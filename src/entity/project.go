package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	alphabet            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ROUND_LIMIT         = 60
	secretLength        = 32
	credentialLegthPart = 32
	keyLegth            = 50
)

type Project struct {
	Name         string             `bson:"name,omitempty"`
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Description  string             `bson:"description,omitempty"`
	HashAlgoritm string             `bson:"hashAlgoritm,omitempty"`
	RoundHash    uint               `bson:"roundHash,omitempty"`
	ApiKey       string             `bson:"apiKey,omitempty"`
	Secret       string             `bson:"secret,omitempty"`
	CreatedBy    string             `bson:"createdBy,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty"`
	UpdatedBy    string             `bson:"updatedBy,omitempty"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty"`
}

func NewProject() *Project {
	return &Project{}
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
