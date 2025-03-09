package app

import (
	"context"
	"fmt"

	"github.com/reinhardlinardi/atm-report/internal/dataset"
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
	// fmt.Println("consumer started")

	for file := range channel {
		info, err := dataset.ParseFileInfo(file)
		if err != nil {
			fmt.Printf("err parse info: %s\n", err.Error())
			continue
		}

		exist, err := app.atmRepo.IsExist(info.AtmId)
		if err != nil {
			fmt.Printf("err check exist atm id: %s\n", info.Name)
			continue
		}
		if !exist {
			fmt.Printf("err atm id not exist: %s\n", info.Name)
			continue
		}

		// raw, err := app.storage.Fetch(file)
		// if err != nil {
		// 	fmt.Printf("err fetch file: %s\n", err.Error())
		// 	continue
		// }

		// fmt.Println(raw)

		// switch ext {
		// // case "csv":
		// // case "json":
		// case "yaml":

		// // case "xml":
		// default:
		// 	fmt.Printf("err unknown format %s: %s\n", ext, file)
		// 	continue
		// }
	}

	// fmt.Println("consumer stopped")
}
