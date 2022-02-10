package usecase

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	DocumentNumber string    `json:"documentNumber"`
	IssuedAt       time.Time `json:"issued_at"`
	ExpiredAt      time.Time `json:"expired_at"`
}

type JWT interface {
	CreateToken(payload Payload) (string, error)

	Verify(token string) (*Payload, error)
}
