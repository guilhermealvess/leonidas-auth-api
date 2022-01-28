package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectRepository interface {
	Insert(project Project) (primitive.ObjectID, error)

	FindByCredentials(credentials string) (*Project, error)

	FindByID(id string) error
}

type AccountRepository interface {
	Insert(name string, desciption string, hashAlgoritm string, roundsHash uint) error

	FindByID(id string) error
}
