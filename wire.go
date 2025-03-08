//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

func initApp(config *config.Config) *app.App {
	wire.Build(app.New, db.New, dbConfig)
	return &app.App{}
}
