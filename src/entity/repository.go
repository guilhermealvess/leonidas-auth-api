package entity

import (
	"github.com/google/uuid"
)

type ProjectRepository interface {
	Insert(project Project) (uuid.UUID, error)

	FindByCredential(credential string) (*Project, error)

	FindByID(id string) error
}

type AccountRepository interface {
	Insert(account Account) (uuid.UUID, error)

	FindByID(id string) error

	FindByEmail(email string, projectId uuid.UUID) (*Account, error)

	Update(account Account) error
}
