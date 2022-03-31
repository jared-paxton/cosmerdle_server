package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func NewApp(cfg Config, logger *log.Logger) (*application, error) {
	database, err := setupDatabase(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	dbAccessor := db.NewDatabaseAccessor(database)

	return &application{
		config:      cfg,
		logger:      logger,
		appServices: newServicers(&dbAccessor),
	}, nil
}

func (app *application) Run() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Println("Starting server on port", app.config.Port)

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func newServicers(dbAccessor *db.DatabaseAccessor) *services {
	return &services{
		gameService: game.NewGameServicer(dbAccessor),
	}
}

func setupDatabase(driver string, connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open(driver, connString)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ping the DB to ensure that it is connected
	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
