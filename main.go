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
		fmt.Printf("err parse config: %s\n", err.Error())
		return
	}

	dbConf := &db.Config{User: conf.DB.User, Pass: conf.DB.Pass,
		Host: conf.DB.Host, Port: conf.DB.Port, Schema: conf.DB.Schema}

	app, err := initApp(conf, dbConf)
	if err != nil {
		fmt.Printf("err init app: %s\n", err.Error())
		return
	}

	if err := app.Connect(); err != nil {
		fmt.Printf("err connect: %s\n", err.Error())
		return
	}
	defer app.Disconnect()

	ctx, cancel := context.WithCancel(context.Background())
	cleanup := make(chan bool, 1)

	go app.Run(ctx, cancel, cleanup)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig
	cancel()

	<-cleanup
}
