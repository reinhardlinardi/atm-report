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
		app.New,
		dbConfig,
		db.New,
		fswatch.New,
	)

	return &app.App{}, nil
}
