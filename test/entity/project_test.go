package entity_test

import (
	"api-auth/src/entity"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidProject(t *testing.T) {
	project := &entity.Project{
		Name:         "Leonidas",
		Description:  "Leonidas Marketplace",
		RoundHash:    8,
		ApiKey:       uuid.NewString(),
		Secret:       "jxzuaiIsNBakqSwQpOQgNczgaczAInLq",
		HashAlgoritm: "sha1",
		CreatedBy:    "SYSTEM",
		CreatedAt:    time.Now(),
	}

	if project.IsValid() != nil {
		t.Errorf("Want not error")
	}
}
