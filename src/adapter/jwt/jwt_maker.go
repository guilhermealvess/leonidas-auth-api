package jwt

import (
	
	"errors"
	"time"

	"github.com/brianvoe/sjwt"
)

type JWTMaker struct {
	JWT
}

func NewJWTMaker() *JWTMaker {
	return &JWTMaker{}
}

func (j *JWTMaker) CreateToken(payload Payload, secret string) (string, error) {
	claims, _ := sjwt.ToClaims(payload)

	secretKey := []byte(secret)
	return claims.Generate(secretKey), nil
}

func (j *JWTMaker) Verify(tokenJWT string, secret string) (*Payload, error) {

	hasVerified := sjwt.Verify(tokenJWT, []byte(secret))
	if !hasVerified {
		return &Payload{}, errors.New("TOKEN INVALIDO")
	}

	payload, _ := j.ParserToken(tokenJWT)

	now := time.Now()
	if now.After(payload.ExpiredAt) {
		return payload, errors.New("TOKEN EXPIRADO")
	}

	return payload, nil
}

func (j *JWTMaker) ParserToken(tokenJWT string) (*Payload, error) {
	claims, err := sjwt.Parse(tokenJWT)
	payload := &Payload{}

	if err == nil {
		claims.ToStruct(payload)
	}

	return payload, err
}
