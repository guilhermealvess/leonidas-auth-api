package jwt

import (
	"fmt"
	"os"

	"github.com/brianvoe/sjwt"
)

type JWTMaker struct {
	JWT
}

func NewJWTMaker() *JWTMaker {
	return &JWTMaker{}
}

func (j *JWTMaker) SignIn() error {
	return nil
}

func (j *JWTMaker) Verify() error {
	return nil
}

func example() {
	// Set Claims
	claims := sjwt.New()
	claims.Set("username", "billymister")
	claims.Set("account_id", 8675309)

	// Generate jwt
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
}
