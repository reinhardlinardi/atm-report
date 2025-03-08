//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/reinhardlinardi/atm-report/app"
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

func initApp(config *app.Config) *app.App {
	wire.Build(app.New, db.New, dbConfig)
	return &app.App{}
}
