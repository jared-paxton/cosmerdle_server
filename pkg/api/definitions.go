package api

import (
	"database/sql"
	"log"

	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

type Config struct {
	Port int
    Database *sql.DB
}

type application struct {
	config      Config
	logger      *log.Logger
	appServices *services
}

type services struct {
	gameService game.GameServicer
}
