package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type App struct {
	db      db.DB
	watcher fswatch.Watcher
	storage storage.Storage
	config  *config.Config
	wg      sync.WaitGroup
}

func New(db db.DB, watcher fswatch.Watcher, storage storage.Storage, config *config.Config) *App {
	return &App{db: db, watcher: watcher, storage: storage, config: config}
}

func (app *App) Run(ctx context.Context, cancel context.CancelFunc, cleanup chan bool) {
	go app.RunCron(ctx, cancel)
	app.RunServer(ctx, cancel)

	// fmt.Println("waiting for goroutines...")
	app.wg.Wait()

	cleanup <- true
}

func (app *App) RunServer(ctx context.Context, cancel context.CancelFunc) {
	addr := fmt.Sprintf(":%d", app.config.Server.Port)
	server := http.Server{
		Addr:    addr,
		Handler: nil,
	}

	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		fmt.Printf("Listening on %s\n", addr)
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) && err != nil {
			fmt.Printf("err server listen: %s", err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	server.Shutdown(context.Background())
	// fmt.Println("server stopped")
}
