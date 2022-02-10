package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectRepository interface {
	Insert(project Project) (primitive.ObjectID, error)

	FindByCredentials(credentials string) (*Project, error)

	FindByID(id string) error
}

type AccountRepository interface {
	Insert(account Account) (primitive.ObjectID, error)

	FindByID(id string) error

	FindByEmail(email string, projectId string) (*Account, error)
}
