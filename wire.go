//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

func initApp(config *config.Config) (*app.App, error) {
	wire.Build(
		dbConfig,
		app.New,
		db.New,
		fswatch.New,
		wire.Bind(new(db.DB), new(*db.DBImpl)),
		wire.Bind(new(fswatch.Watcher), new(*fswatch.WatcherImpl)),
	)

	return &app.App{}, nil
}
