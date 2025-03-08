package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/db"
)

func main() {
	dir := "config"
	file := "config.yaml"

	config, err := app.ParseConfig(path.Join(dir, file))
	if err != nil {
		return
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	initApp(config)
	fmt.Println("DB connected")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("shutting down")
}

func dbConfig(config *app.Config) *db.Config {
	return &db.Config{
		User:   config.DB.User,
		Pass:   config.DB.Pass,
		Host:   config.DB.Host,
		Port:   config.DB.Port,
		Schema: config.DB.Schema,
	}
}
