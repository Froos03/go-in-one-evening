package account

import "errors"

var (
	ErrEmptyEmail    = errors.New("empty email")
	ErrEmptyPassword = errors.New("empty password")
)

type Account struct {
	Email    string
	Password string
}

func (a Account) Login(email string, password string) bool {
	return a.Email == email && a.Password == password
}

func New(email string, password string) (Account, error) {
	if email == "" {
		return Account{}, ErrEmptyEmail
	}

	if password == "" {
		return Account{}, ErrEmptyPassword
	}

	return Account{
		Email:    email,
		Password: password,
	}, nil
}
