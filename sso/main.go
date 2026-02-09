package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"sso/application"
	"sso/domain"
	"sso/infrastructure/repository"
	"sso/infrastructure/rest"
)

func main() {
	fmt.Println("Server started")

	ctx := context.TODO()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:example@db:27017/"))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to MongoDB!")

	repositories := initRepositories(&ctx, client)

	var app, cleanup = application.NewApplication(repositories)
	defer cleanup()

	router := rest.NewRouter(mux.NewRouter(), app)

	err = router.Serve(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func initRepositories(ctx *context.Context, dbClient *mongo.Client) domain.Repositories {
	userRepo := repository.NewUserRepository(ctx, dbClient)

	return domain.Repositories{
		UserRepository: userRepo,
	}
}
