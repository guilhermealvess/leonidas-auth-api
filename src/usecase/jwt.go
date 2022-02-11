package usecase

import (
	"time"
)

type Payload struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type JWT interface {
	CreateToken(payload Payload) (string, error)

	Verify(token string) (*Payload, error)
}
