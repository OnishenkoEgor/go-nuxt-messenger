package command

import (
	"github.com/google/uuid"
	"sso/domain/user"
)

type DeleteUserCommand struct {
	Id uuid.UUID
}

type DeleteUserCommandHandler struct {
	repo user.Repository
}

func NewDeleteUserCommandHandler(repo user.Repository) DeleteUserCommandHandler {
	return DeleteUserCommandHandler{
		repo: repo,
	}
}

func (h DeleteUserCommandHandler) Handle(cmd DeleteUserCommand) error {
	err := h.repo.Delete(cmd.Id.String())

	return err
}
