package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

func main() {
	conf, err := config.Parse(path.Join("config", "config.yaml"))
	if err != nil {
		return
	}

	app, err := initApp(conf)
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	cleanup := make(chan bool, 1)

	go app.Run(ctx, cancel, cleanup)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	fmt.Println("\ninterrupted, shutting down...")
	cancel()

	<-cleanup
	fmt.Println("exited")
}

func dbConfig(conf *config.Config) *db.Config {
	return &db.Config{
		User:   conf.DB.User,
		Pass:   conf.DB.Pass,
		Host:   conf.DB.Host,
		Port:   conf.DB.Port,
		Schema: conf.DB.Schema,
	}
}
