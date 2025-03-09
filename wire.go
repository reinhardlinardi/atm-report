//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/atmrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/historyrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
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
		atmrepo.New,
		wire.Bind(new(atmrepo.Repository), new(*atmrepo.RepositoryImpl)),
		historyrepo.New,
		wire.Bind(new(historyrepo.Repository), new(*historyrepo.RepositoryImpl)),
		transactionrepo.New,
		wire.Bind(new(transactionrepo.Repository), new(*transactionrepo.RepositoryImpl)),
	)

	return &app.App{}, nil
}
