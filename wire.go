//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/app/cron"
	"github.com/reinhardlinardi/atm-report/app/server"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/history"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

func initApp(conf *config.Config, dbConf *db.Config) (*app.App, error) {
	wire.Build(
		app.New,
		server.New,
		cron.New,
		wire.FieldsOf(new(*config.Config), "Server"),
		wire.FieldsOf(new(*config.Config), "Cron"),
		db.New,
		wire.Bind(new(db.DB), new(*db.DBImpl)),
		fswatch.New,
		wire.Bind(new(fswatch.Watcher), new(*fswatch.WatcherImpl)),
		storage.New,
		wire.Bind(new(storage.Storage), new(*storage.StorageImpl)),
		history.New,
		wire.Bind(new(history.Repository), new(*history.RepositoryImpl)),
		transaction.New,
		wire.Bind(new(transaction.Repository), new(*transaction.RepositoryImpl)),
	)

	return &app.App{}, nil
}
