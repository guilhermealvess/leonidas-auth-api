package entity

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const PASSWORD_LENGTH = 8
const USERNAME_LENGTH = 8

type Account struct {
	ID            string    `json:"id"`
	ProjectID     string    `json:"projectID"`
	UID           uuid.UUID `json:"uid"`
	FirstName     string    `json:"firtName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	VerifiedEmail bool      `json:"verifiedEmail"`
	IsActive      bool      `json:"isActive"`
	ActivedAt     time.Time `json:"activedAt"`
	LastLogin     time.Time `json:"lastLogin"`
}

func NewAccount() *Account {
	uid, _ := uuid.NewUUID()
	return &Account{
		UID: uid,
	}
}

func (a *Account) IsValid() error {
	if a.FirstName == "" || a.LastName == "" {
		return errors.New("Name and LastName is required")
	}

	if len(a.Password) < PASSWORD_LENGTH {
		return errors.New("Password invalid")
	}

	if len(a.ProjectID) == 0 {
		return errors.New("ProjectId is required")
	}

	err := a.ValidEmail(a.Email)
	if err != nil {
		return err
	}

	if len(a.Username) < USERNAME_LENGTH {
		return errors.New("Username invalid")
	}

	return nil
}

func (a *Account) ValidEmail(email string) error {
	validate := validator.New()
	return validate.Var(email, "email")
}

func (a *Account) SavePassword(password, algorithm string, rounds uint) error {
	algorithm = strings.ToUpper(algorithm)
	if algorithm != "SHA256" && algorithm != "SHA1" && algorithm != "SHA512" && algorithm != "MD5" {
		return errors.New("Algorithm invalid")
	}
	for i := 0; i < int(rounds); i++ {
		password = calculateHash(password, algorithm)
	}

	a.Password = password
	return nil
}

func calculateHash(str string, algorithm string) string {
	strHash := str
	switch strings.ToUpper(algorithm) {
	case "SHA256":
		hash := sha256.New()
		hash.Write([]byte(str))
		strHash = hex.EncodeToString(hash.Sum(nil))

	case "SHA1":
		hash := sha1.New()
		hash.Write([]byte(str))
		strHash = hex.EncodeToString(hash.Sum(nil))

	case "SHA512":
		hash := sha512.New()
		hash.Write([]byte(str))
		strHash = hex.EncodeToString(hash.Sum(nil))

	case "MD5":
		hash := md5.New()
		hash.Write([]byte(str))
		strHash = hex.EncodeToString(hash.Sum(nil))
	}

	return strHash
}

func (a *Account) VerifyPassword(password string, rounds uint, algorithm string) bool {
	for i := 0; i < int(rounds); i++ {
		password = calculateHash(password, algorithm)
	}

	return password == a.Password
}

func (a *Account) ActivedAccount() {
	a.ActivedAt = time.Now()
	a.IsActive = true
}
