package entity_test

import (
	"api-auth/src/entity"
	"testing"
)

func TestValidEmail(t *testing.T) {
	account := &entity.Account{}

	if account.ValidEmail("guilherme.alves@gmail.com") != nil {
		t.Errorf("Want not error")
	}

	//TODO
	/*if account.ValidEmail("@gmail.com") == nil || account.ValidEmail("ofadPJASODJ.com") == nil {
		t.Errorf("Want error")
	}*/

}

func TestSavePassword(t *testing.T) {
	account := &entity.Account{}
	algorithm := "sha1"

	if account.SavePassword("Teste#123", algorithm, uint(5)) == nil {
		if account.Password != "fdb1d4219312d178f84c51240950f84572f8dee4" {
			t.Errorf("Hash calclute invalid")
		}
	} else {
		t.Errorf("Want not error")
	}
}

func TestVerifyPassword(t *testing.T) {
	account := &entity.Account{}
	round := uint(5)
	algorithm := "sha1"
	password := "Teste#123a"

	account.SavePassword(password, algorithm, round)
	if !account.VerifyPassword(password, round, algorithm) {
		t.Errorf("Expected return true")
	}
}
