package command

import (
	"messenger/sso/domain/user"
)

type UpdateUserCommand struct {
	Id       int
	Login    string
	Password string
}

type UpdateUserCommandHandler struct {
	repo user.Repository
}

func NewUpdateUserCommandHandler(repo user.Repository) UpdateUserCommandHandler {
	return UpdateUserCommandHandler{
		repo: repo,
	}
}

func (r *UpdateUserCommandHandler) Handle(cmd UpdateUserCommand) error {
	u, err := r.repo.GetById(cmd.Id)
	if err != nil {
		return err
	}

	u.SetLogin(cmd.Login)
	u.SetPassword(cmd.Password)

	return r.repo.Update(u)
}
