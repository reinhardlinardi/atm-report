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
	db      db.DB
	watcher fswatch.Watcher
	config  *config.Config
}

func New(db db.DB, watcher fswatch.Watcher, config *config.Config) *App {
	return &App{db: db, watcher: watcher, config: config}
}

func (app *App) Run(ctx context.Context, cleanup chan bool) {
	var wg sync.WaitGroup

	wg.Add(1)
	files := make(chan string, 26)

	go func() {
		defer wg.Done()
		app.watcher.WatchCreated(ctx, "dataset", files)
	}()

	wg.Wait()
	fmt.Println("wait complete")

	cleanup <- true
}
