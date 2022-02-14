package entity

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ProjectId primitive.ObjectID `bson:"projectId,omitempty"`
	Name      string             `bson:"name,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	LastLogin string             `bson:"lastLogin,omitempty"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) IsValid() error {
	if a.Name == "" || a.LastName == "" {
		return errors.New("Name and LastName is required")
	}

	if len(a.Password) < 8 {
		return errors.New("Password invalid")
	}

	if len(a.ProjectId) == 0 {
		return errors.New("ProjectId is required")
	}

	err := a.ValidEmail(a.Email)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) ValidEmail(email string) error {
	valid := regexp.MustCompile(`/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/`)

	if valid.MatchString(email) {
		return errors.New("Email invalido")
	}

	return nil
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
