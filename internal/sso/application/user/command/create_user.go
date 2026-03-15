package command

import (
	"messenger/sso/domain/user"
)

type CreateUserCommand struct {
	Login    string
	Password string
}

type CreateUserCommandHandler struct {
	repo user.Repository
}

func NewCreateUserCommandHandler(repo user.Repository) CreateUserCommandHandler {
	return CreateUserCommandHandler{
		repo: repo,
	}
}

func (h CreateUserCommandHandler) Handle(cmd CreateUserCommand) error {
	u, err := user.NewUser(cmd.Login, cmd.Password)

	if err != nil {
		return err
	}

	err = h.repo.Create(u)

	if err != nil {
		return err
	}

	return nil
}
