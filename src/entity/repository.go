package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectRepository interface {
	Insert(project Project) (primitive.ObjectID, error)

	FindByApiKey(apiKey string) (*Project, error)

	FindByID(id string) (*Project, error)
}

type AccountRepository interface {
	Insert(account Account) (primitive.ObjectID, error)

	FindByID(id string) (*Account, error)

	FindByEmail(email string, projectId primitive.ObjectID) (*Account, error)

	UpdateActived(id primitive.ObjectID) error
}
