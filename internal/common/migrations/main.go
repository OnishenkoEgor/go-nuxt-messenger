package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	dbUrl := "postgres://user:password@db/sso?sslmode=verify-full"
	dbName := "postgres"
	db, err := sql.Open(dbName, dbUrl)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(os.Args))
	//if len(os.Args) < 2{
	//
	//}
	//command := os.Args[1];

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println(os.Args[1])
}

type Migrations struct {
	db *sql.DB
}

func (m *Migrations) Migrate() error {
	//t, err := m.db.Begin()
	//if err != nil {
	//	return err
	//}
	//
	//for _, migration := range m.migrationsList {
	//	_, err = m.db.Exec(migration.upCommand)
	//	if err != nil {
	//		fmt.Println("Migration failed", err)
	//
	//		err = t.Rollback()
	//
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}
	//
	//err = t.Commit()
	//if err != nil {
	//	return err
	//}

	return nil
}

func (m *Migrations) Rollback() error {
	return nil
}
