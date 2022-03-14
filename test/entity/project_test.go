package entity_test

import (
	"api-auth/src/entity"
	"testing"
	"time"
)

func TestValidProject(t *testing.T) {
	project := &entity.Project{
		Name: "Leonidas",
		Description: "Leonidas Marketplace",
		RoundHash: 8,
		Key: "255e5f3f-c8d0-5c96-8cf5-3eac80facb60",
		Credential: "fgPyCKmxIubeYTNDtjAyRRDedMiyLpru-cjiOgjhYeVwBTCMLfrDGXqwpzwVGqMZc",
		Secret: "jxzuaiIsNBakqSwQpOQgNczgaczAInLq",
		HashAlgoritm: "sha1",
		CreatedBy: "SYSTEM",
		CreatedAt: time.Now(),
	}

	if project.IsValid() != nil {
		t.Errorf("Want not error")
	}
}
