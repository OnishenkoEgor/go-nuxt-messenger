package user

import (
	"errors"
	"github.com/google/uuid"
)

func NewUser(uuid uuid.UUID, login string, password string) (*User, error) {
	if len(login) == 0 {
		return nil, errors.New("login cannot be empty")
	}

	if len(password) == 0 {
		return nil, errors.New("password cannot be empty")
	}

	return &User{
		uuid,
		login,
		password,
	}, nil
}

type User struct {
	uuid     uuid.UUID
	login    string
	password string
}

func (user *User) GetUUID() uuid.UUID {
	return user.uuid
}

func (user *User) GetLogin() string {
	return user.login
}

func (user *User) GetPassword() string {
	return user.password
}
