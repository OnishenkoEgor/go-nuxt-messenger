package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"messenger/router"
	"messenger/sso/application"
	"messenger/sso/infrastructure"
)

func main() {
	fmt.Println("Server started")

	//TODO ssl?
	dbUrl := "postgres://user:password@db/sso?sslmode=disable"
	dbName := "postgres"
	db, err := sql.Open(dbName, dbUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to DB!")

	repositories := infrastructure.NewRepositories(db)
	app, cleanup := application.NewApplication(repositories)
	defer cleanup()

	routes := infrastructure.NewRoutes(app)
	r, err := router.NewRouter(routes)

	if err != nil {
		fmt.Println(err)
	}

	err = r.Serve(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
