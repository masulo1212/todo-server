package main

import (
	"database/sql"
	"log"
	"os"
	"todo/broker/api"

	// "github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	DB_driver = "postgres"
	// DB_source    = "postgresql://root:secret@localhost:5432/todoDB?sslmode=disable"
	// migrationUrl = "file://db/migration"
)

func main() {

	db_source := os.Getenv("DSN")
	db, err := sql.Open(DB_driver, db_source)
	if err != nil {
		log.Fatalln("cannot connect to db")
		return
	}

	// migration, err := migrate.New(migrationUrl, db_source)
	// if err != nil {
	// 	log.Fatalf("cannot create new migration instance: %d", err)
	// 	return
	// }

	// if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("failed to run migrate up: %d", err)
	// 	return
	// }

	// log.Println("db migrated successfully")

	server := api.NewServer(db)

	log.Println("server run at :9090...")
	if err = server.Start(); err != nil {
		log.Fatalln("cannot connect to :9090")
		return
	}

}
