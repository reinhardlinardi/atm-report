package app

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/reinhardlinardi/atm-report/internal/config"
	"github.com/reinhardlinardi/atm-report/internal/repository/atmrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/historyrepo"
	"github.com/reinhardlinardi/atm-report/internal/repository/transactionrepo"
	"github.com/reinhardlinardi/atm-report/internal/storage"
	"github.com/reinhardlinardi/atm-report/pkg/fswatch"
)

type Cron struct {
	config          *config.Config
	watcher         fswatch.Watcher
	storage         storage.Storage
	atmRepo         atmrepo.Repository
	historyRepo     historyrepo.Repository
	transactionRepo transactionrepo.Repository
}

func NewCron(
	config *config.Config,
	watcher fswatch.Watcher,
	storage storage.Storage,
	atmRepo atmrepo.Repository,
	historyRepo historyrepo.Repository,
	transactionRepo transactionrepo.Repository,
) *Cron {
	return &Cron{
		config:          config,
		watcher:         watcher,
		storage:         storage,
		atmRepo:         atmRepo,
		historyRepo:     historyRepo,
		transactionRepo: transactionRepo,
	}
}

func (c *Cron) Run(ctx context.Context, cancel context.CancelFunc) {
	channel := make(chan string, 10)

	go c.runWatcher(ctx, cancel, channel)
	c.runConsumer(ctx, cancel, channel)
}

func (c *Cron) runWatcher(ctx context.Context, cancel context.CancelFunc, channel chan string) {
	err := c.watcher.WatchCreated(ctx, c.config.Cron.Path, channel)

	if err != nil {
		fmt.Printf("err watcher: %s\n", err.Error())
		cancel()
		close(channel)
	}
}

func (c *Cron) runConsumer(_ context.Context, _ context.CancelFunc, channel chan string) {
	for path := range channel {
		if err := c.handleFile(path); err != nil {
			fmt.Printf("err handle file: %s\n", err.Error())
		} else {
			fmt.Printf("%s processed\n", filepath.Base(path))
		}
	}
}
