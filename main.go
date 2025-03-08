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
	dir := "config"
	file := "config.yaml"

	conf, err := config.Parse(path.Join(dir, file))
	if err != nil {
		return
	}

	app := initApp(conf)
	ctx, cancel := context.WithCancel(context.Background())
	shutdown := make(chan bool, 1)

	go app.Run(ctx, shutdown)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	fmt.Println("waiting...")

	<-sig
	fmt.Println("shutting down")
	cancel()

	<-shutdown
	fmt.Println("shutdown complete")
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
