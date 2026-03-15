package application

import (
	"fmt"
	"messenger/sso/application/user/command"
	"messenger/sso/application/user/query"
	"messenger/sso/domain"
)

type Queries struct {
	GetUsersQuery    query.GetUsersQueryHandler
	GetUserByIdQuery query.GetUserByIdQueryHandler
}

type Commands struct {
	CreateUserCommand command.CreateUserCommandHandler
	DeleteUserCommand command.DeleteUserCommandHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(repositories domain.Repositories) (Application, func()) {
	var app = Application{
		Commands: Commands{
			CreateUserCommand: command.NewCreateUserCommandHandler(repositories.UserRepository),
			DeleteUserCommand: command.NewDeleteUserCommandHandler(repositories.UserRepository),
		},
		Queries: Queries{
			GetUsersQuery:    query.NewGetUsersQueryHandler(repositories.UserRepository),
			GetUserByIdQuery: query.NewGetUserByIdQueryHandler(repositories.UserRepository),
		},
	}

	return app, func() {
		fmt.Println("Cleanup...")
	}
}
