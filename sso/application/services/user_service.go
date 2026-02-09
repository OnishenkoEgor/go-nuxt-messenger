package services

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sso/domain/user"
	"sso/infrastructure/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(ctx *context.Context, dbClient *mongo.Client) UserService {
	userRepository := repository.NewUserRepository(ctx, dbClient)

	return UserService{repository: userRepository}
}

func (service UserService) GetById(id string) (*user.User, error) {

	return service.repository.GetById(id)
}

func (service UserService) GetAll() []*user.User {
	return service.repository.Get()
}

func (service UserService) Create(login string, password string) error {
	id, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	user, err := user.NewUser(id, login, password)

	if err != nil {
		return err
	}

	err = service.repository.Create(user)

	return err
}

func (service UserService) Delete(id string) error {
	return service.repository.Delete(id)
}
