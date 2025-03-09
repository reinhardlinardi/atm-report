package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/atmrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/historyrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	"github.com/reinhardlinardi/atm-report/pkg/db"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type App struct {
	db              db.DB
	watcher         fswatch.Watcher
	storage         storage.Storage
	atmRepo         atmrepo.Repository
	historyRepo     historyrepo.Repository
	transactionRepo transactionrepo.Repository
	config          *config.Config
	wg              sync.WaitGroup
}

func New(
	config *config.Config,
	db db.DB,
	watcher fswatch.Watcher,
	storage storage.Storage,
	atmRepo atmrepo.Repository,
	historyRepo historyrepo.Repository,
	transactionRepo transactionrepo.Repository,
) *App {
	return &App{
		db:              db,
		watcher:         watcher,
		storage:         storage,
		atmRepo:         atmRepo,
		historyRepo:     historyRepo,
		transactionRepo: transactionRepo,
		config:          config,
	}
}

func (app *App) Connect() error {
	if err := app.db.Connect(); err != nil {
		return fmt.Errorf("err connect db: %s", err.Error())
	}
	return nil
}

func (app *App) Disconnect() {
	app.db.Disconnect()
}

func (app *App) Run(ctx context.Context, cancel context.CancelFunc, cleanup chan bool) {
	go app.runCron(ctx, cancel)
	app.runServer(ctx, cancel)

	app.wg.Wait()
	cleanup <- true
}

func (app *App) runServer(ctx context.Context, cancel context.CancelFunc) {
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
			fmt.Printf("err server listen: %s\n", err.Error())
			cancel()
		}
	}()

	<-ctx.Done()
	server.Shutdown(context.Background())
}
