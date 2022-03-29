package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jared-paxton/cosmerdle_server/pkg/api"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Cosmerdle!")

	connectionString := "postgres://postgres:postgres@localhost/cosmerdle?sslmode=disable"

	database, err := setupDatabase(connectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := api.Config{
		Port:     4001,
		Database: database,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := api.NewApp(cfg, logger)
	err = app.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
