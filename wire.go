//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/atmrepository"
	"github.com/reinhardlinardi/atm-report/internal/repository/fileloadrepository"
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
		atmrepository.New,
		wire.Bind(new(atmrepository.Repository), new(*atmrepository.RepositoryImpl)),
		fileloadrepository.New,
		wire.Bind(new(fileloadrepository.Repository), new(*fileloadrepository.RepositoryImpl)),
	)

	return &app.App{}, nil
}
