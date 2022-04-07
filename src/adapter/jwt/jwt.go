package jwt

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
	CreateToken(payload Payload, secret string) (string, error)

	Verify(token string, secret string) (*Payload, error)

	ParserToken(tokenJWT string) (*Payload, error)
}
