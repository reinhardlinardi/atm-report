package app

import (
	"github.com/reinhardlinardi/atm-report/internal/db"
)

type App struct {
	DB *db.DB
}

func New(db *db.DB) *App {
	return &App{DB: db}
}
