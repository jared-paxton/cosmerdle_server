package app

import (
	"log"

	"github.com/jared-paxton/cosmerdle_server/pkg/game"
	"github.com/jared-paxton/cosmerdle_server/pkg/user"
)

type Config struct {
	Port     int
	Env      string
	Database struct {
		DSN    string
		Driver string
	}
}

type application struct {
	config      Config
	logger      *log.Logger
	appServices *services
}

type services struct {
	gameService game.GameServicer
	userService user.UserServicer
}
