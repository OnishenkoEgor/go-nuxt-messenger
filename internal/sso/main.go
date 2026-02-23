package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"github.com/gorilla/mux"
	"sso/application"
	"sso/infrastructure"
)

func main() {
	fmt.Println("Server started")

	dbUrl := "postgres://user:password@db/sso?sslmode=verify-full"
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

	if false {
		migrations := infrastructure.NewMigrations(db)
		err = migrations.Migrate()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("DB migrations success!")
	}

	repositories := infrastructure.NewRepositories(db)
	app, cleanup := application.NewApplication(repositories)
	defer cleanup()

	router := infrastructure.NewRouter(mux.NewRouter(), app)
	err = router.Serve(":8080")

	if err != nil {
		fmt.Println(err)
	}
}
