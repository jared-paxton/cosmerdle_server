package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func newServicers(dbAccessor *db.DatabaseAccessor) *services {
	return &services{
		gameService: game.NewGameServicer(dbAccessor),
	}
}

func NewApp(cfg Config, logger *log.Logger) *application {
	dbAccessor := db.NewDatabaseAccessor(cfg.Database)

	return &application{
		config:      cfg,
		logger:      logger,
		appServices: newServicers(&dbAccessor),
	}
}

func (app *application) Run() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
