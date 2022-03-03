package jwt

import (
	"api-auth/src/usecase"
	"errors"
	"time"

	"github.com/brianvoe/sjwt"
)

type JWTMaker struct {
	usecase.JWT
}

func NewJWTMaker() *JWTMaker {
	return &JWTMaker{}
}

func (j *JWTMaker) CreateToken(payload usecase.Payload, secret string) (string, error) {
	claims, _ := sjwt.ToClaims(payload)

	secretKey := []byte(secret)
	return claims.Generate(secretKey), nil
}

func (j *JWTMaker) Verify(tokenJWT string, secret string) (*usecase.Payload, error) {

	hasVerified := sjwt.Verify(tokenJWT, []byte(secret))
	if !hasVerified {
		return &usecase.Payload{}, errors.New("TOKEN INVALIDO")
	}

	claims, _ := sjwt.Parse(tokenJWT)
	payload := &usecase.Payload{}
	claims.ToStruct(payload)

	now := time.Now()
	if now.After(payload.ExpiredAt) {
		return payload, errors.New("TOKEN EXPIRADO")
	}

	return payload, nil
}
