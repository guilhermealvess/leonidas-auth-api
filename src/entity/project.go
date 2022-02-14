package entity

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Project struct {
	Name         string             `bson:"name,omitempty"`
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Description  string             `bson:"description,omitempty"`
	HashAlgoritm string             `bson:"hashAlgoritm,omitempty"`
	RoundHash    uint               `bson:"roundHash,omitempty"`
	Credential   string             `bson:"credential,omitempty"`
	Key          string             `bson:"key,omitempty"`
	Secret       string             `bson:"secret,omitempty"`
	CreatedBy    string             `bson:"createdBy,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty"`
	UpdatedBy    string             `bson:"updatedBy,omitempty"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty"`
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) GenerateKey() string {
	uid, _ := uuid.NewUUID()
	keyLegth := 50
	key := uuid.NewSHA1(uid, []byte(p.generateStringRandom(uint(keyLegth))))
	return key.String()
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
	if p.RoundHash < 1 || p.RoundHash > 60 {
		return errors.New("Round Hash over limit permit")
	}

	if p.Description == "" {
		return errors.New("Description invalid")
	}

	p.Description = strings.ToLower(p.Description)

	return nil
}
