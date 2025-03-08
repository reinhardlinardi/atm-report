package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type App struct {
	db     *db.DB
	config *config.Config
}

func New(db *db.DB, config *config.Config) *App {
	return &App{db: db, config: config}
}

func (app *App) Run(ctx context.Context, shutdown chan bool) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fswatch.Run(ctx, "dataset")
	}()

	wg.Wait()
	fmt.Println("wait complete")

	shutdown <- true
}
