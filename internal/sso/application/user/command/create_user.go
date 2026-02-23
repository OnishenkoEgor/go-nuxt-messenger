package command

import (
	"github.com/google/uuid"
	"sso/domain/user"
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
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u, err := user.NewUser(uuid.String(), cmd.Login, cmd.Password)
	if err != nil {
		return err
	}

	err = h.repo.Create(u)
	if err != nil {
		return err
	}

	return nil
}
