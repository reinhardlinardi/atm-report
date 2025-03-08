package app

import (
	"context"
	"fmt"
)

func (app *App) RunCron(ctx context.Context, cancel context.CancelFunc) {
	files := make(chan string, 10)

	app.wg.Add(2)
	defer app.wg.Done()

	go func() {
		defer app.wg.Done()

		err := app.watcher.WatchCreated(ctx, app.config.Cron.Path, files)
		if err != nil {
			fmt.Printf("err watcher: %s\n", err.Error())
			close(files)
			cancel()
		}
	}()

	fmt.Println("cron started")

	for file := range files {
		fmt.Println(file)
	}

	fmt.Println("cron stopped")
}
