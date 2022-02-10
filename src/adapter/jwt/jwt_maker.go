package jwt

import (
	"api-auth/src/usecase"
	"errors"
	"os"
	"time"

	"github.com/brianvoe/sjwt"
)

type JWTMaker struct {
	usecase.JWT
}

func NewJWTMaker() *JWTMaker {
	return &JWTMaker{}
}

func (j *JWTMaker) CreateToken(payload usecase.Payload) (string, error) {
	claims, _ := sjwt.ToClaims(payload)

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	return claims.Generate(secretKey), nil
}

func (j *JWTMaker) Verify(tokenJWT string) (*usecase.Payload, error) {

	hasVerified := sjwt.Verify(tokenJWT, []byte(os.Getenv("JWT_SECRET")))
	if !hasVerified {
		return &usecase.Payload{}, errors.New("TOKEN INVALIDO")
	}

	claims, _ := sjwt.Parse(tokenJWT)
	payload := &usecase.Payload{}
	claims.ToStruct(payload)

	if time.Now().After(payload.ExpiredAt) {
		return payload, errors.New("TOKEN EXPIRADO")
	}

	return payload, nil
}
