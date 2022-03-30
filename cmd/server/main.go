package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jared-paxton/cosmerdle_server/pkg/app"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Cosmerdle!")

	var cfg app.Config

	flag.IntVar(&cfg.Port, "Port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.Env, "Env", "development", "Application environment (development|production")
	flag.StringVar(&cfg.Database.DSN, "DSN", "postgres://jaredpaxton@localhost/cosmerdle?sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.Database.Driver, "Driver", "postgres", "Postgres connection string")
	//flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Println(err)
		return
	}

	err = app.Run()
	if err != nil {
		logger.Println(err)
	}
}
