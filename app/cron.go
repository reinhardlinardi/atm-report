package app

import (
	"context"
	"fmt"
	"path/filepath"
)

func (app *App) runCron(ctx context.Context, cancel context.CancelFunc) {
	app.wg.Add(2)
	channel := make(chan string, 10)

	go app.runWatcher(ctx, cancel, channel)
	app.runConsumer(ctx, cancel, channel)
}

func (app *App) runWatcher(ctx context.Context, cancel context.CancelFunc, channel chan string) {
	defer app.wg.Done()
	err := app.watcher.WatchCreated(ctx, app.config.Cron.Path, channel)

	if err != nil {
		fmt.Printf("err watcher: %s\n", err.Error())
		close(channel)
		cancel()
	}
}

func (app *App) runConsumer(_ context.Context, _ context.CancelFunc, channel chan string) {
	defer app.wg.Done()

	for path := range channel {
		if err := app.handleFile(path); err != nil {
			fmt.Printf("err handle file: %s\n", err.Error())
		} else {
			fmt.Printf("%s processed\n", filepath.Base(path))
		}
	}
}
