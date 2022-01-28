package jwt

type JWT interface {
	SignIn() error

	Verify() error
}
