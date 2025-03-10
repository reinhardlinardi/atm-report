package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/reinhardlinardi/atm-report/app/cron"
	"github.com/reinhardlinardi/atm-report/app/server"
	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/pkg/db"
)

type App struct {
	config *config.Config
	wg     sync.WaitGroup
	cron   *cron.Cron
	server *server.Server
	db     db.DB
}

func New(
	config *config.Config,
	cron *cron.Cron,
	server *server.Server,
	db db.DB,
) *App {
	return &App{
		config: config,
		cron:   cron,
		server: server,
		db:     db,
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

func (app *App) runCron(ctx context.Context, cancel context.CancelFunc) {
	app.wg.Add(1)
	defer app.wg.Done()

	app.cron.Run(ctx, cancel)
}

func (app *App) runServer(ctx context.Context, cancel context.CancelFunc) {
	app.wg.Add(1)
	defer app.wg.Done()

	app.server.RegisterHandlers()
	go app.server.Run(ctx, cancel)

	<-ctx.Done()
	app.server.Shutdown(context.Background())
}
