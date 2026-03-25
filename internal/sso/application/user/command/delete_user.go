package command

import (
	"messenger/sso/domain/user"
)

type DeleteUserCommand struct {
	Id int
}

type DeleteUserCommandHandler struct {
	repo user.Repository
}

func NewDeleteUserCommandHandler(repo user.Repository) DeleteUserCommandHandler {
	return DeleteUserCommandHandler{
		repo: repo,
	}
}

func (h *DeleteUserCommandHandler) Handle(cmd DeleteUserCommand) error {
	err := h.repo.Delete(cmd.Id)

	return err
}
