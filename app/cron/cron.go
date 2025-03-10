package cron

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/reinhardlinardi/atm-report/internal/config"
	historydb "github.com/reinhardlinardi/atm-report/internal/history"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	transactiondb "github.com/reinhardlinardi/atm-report/internal/transaction"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type Cron struct {
	config        *config.CronConfig
	watcher       fswatch.Watcher
	storage       storage.Storage
	historyDB     historydb.Repository
	transactionDB transactiondb.Repository
}

func New(
	config *config.CronConfig,
	watcher fswatch.Watcher,
	storage storage.Storage,
	historyDB historydb.Repository,
	transactionDB transactiondb.Repository,
) *Cron {
	return &Cron{
		config:        config,
		watcher:       watcher,
		storage:       storage,
		historyDB:     historyDB,
		transactionDB: transactionDB,
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
		fmt.Printf("err watcher: %s\n", err.Error())
		cancel()
		close(channel)
	}
}

func (cron *Cron) runConsumer(_ context.Context, _ context.CancelFunc, channel chan string) {
	for path := range channel {
		if err := cron.handleFile(path); err != nil {
			fmt.Printf("err handle file: %s\n", err.Error())
		} else {
			fmt.Printf("finished: %s\n", filepath.Base(path))
		}
	}
}
