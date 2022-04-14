package entity

type ProjectRepository interface {
	Insert(project Project) (string, error)

	FindByApiKey(apiKey string) (*Project, error)

	FindByID(id string) (*Project, error)
}

type AccountRepository interface {
	Insert(account Account) (string, error)

	FindByID(id string) (*Account, error)

	FindByUsernameAndProject(username string, projectID string) (*Account, error)

	UpdateActived(id string) error

	UpdateLastLogin(id string) error

	UpdatePassword(id string, password string) error
}
	