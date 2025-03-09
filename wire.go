//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

func initApp(config *config.Config) (*app.App, error) {
	wire.Build(
		app.New,
		db.New,
		dbConfig,
		wire.Bind(new(db.DB), new(*db.DBImpl)),
		fswatch.New,
		wire.Bind(new(fswatch.Watcher), new(*fswatch.WatcherImpl)),
		storage.New,
		wire.Bind(new(storage.Storage), new(*storage.StorageImpl)),
	)

	return &app.App{}, nil
}
