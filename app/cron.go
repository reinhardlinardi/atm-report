package app

import (
	"context"
	"fmt"
)

func (app *App) RunCron(ctx context.Context, cancel context.CancelFunc) {
	app.wg.Add(2)
	files := make(chan string, 10)

	go app.RunWatcher(ctx, cancel, files)
	defer app.wg.Done()

	fmt.Println("cron started")

	for file := range files {
		fmt.Println(file)
	}

	fmt.Println("cron stopped")
}

func (app *App) RunWatcher(ctx context.Context, cancel context.CancelFunc, channel chan string) {
	defer app.wg.Done()
	err := app.watcher.WatchCreated(ctx, app.config.Cron.Path, channel)

	if err != nil {
		fmt.Printf("err watcher: %s\n", err.Error())
		close(channel)
		cancel()
	}
}
