package cron

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/filestorage"
	"github.com/reinhardlinardi/atm-report/internal/history"
	"github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type Cron struct {
	config      *config.CronConfig
	watcher     fswatch.Watcher
	fileStorage filestorage.Storage
	history     history.Repository
	transaction transaction.Repository
}

func New(
	config *config.CronConfig,
	watcher fswatch.Watcher,
	fileStorage filestorage.Storage,
	history history.Repository,
	transaction transaction.Repository,
) *Cron {
	return &Cron{
		config:      config,
		watcher:     watcher,
		fileStorage: fileStorage,
		history:     history,
		transaction: transaction,
	}
}

func (cron *Cron) Run(ctx context.Context, cancel context.CancelFunc) {
	channel := make(chan string, 10)

	go cron.runWatcher(ctx, cancel, channel)
	cron.runConsumer(ctx, cancel, channel)
}

func (cron *Cron) runWatcher(ctx context.Context, cancel context.CancelFunc, channel chan string) {
	err := cron.watcher.WatchCreated(ctx, cron.config.Path, channel)

	if err != nil {
		fmt.Printf("err start watcher: %s\n", err.Error())
		cancel()
		close(channel)
	}
}

func (cron *Cron) runConsumer(_ context.Context, _ context.CancelFunc, channel chan string) {
	for path := range channel {
		filename := filepath.Base(path)
		processed, err := cron.handleFile(path)

		if err != nil {
			fmt.Printf("err handle file: %s\n", err.Error())
			continue
		}

		if processed {
			fmt.Printf("done: %s\n", filename)
		} else {
			fmt.Printf("skipped: %s\n", filename)
		}
	}
}
