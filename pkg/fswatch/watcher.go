package fswatch

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func Run(ctx context.Context, path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := watcher.Add(path); err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			switch {
			// useful for new files when watching directories
			case event.Op&fsnotify.Create == fsnotify.Create:
				fmt.Printf("File has been created: %s\n", event.Name)
			}
		case <-ctx.Done():
			fmt.Println("quit watching")
			watcher.Close()
			fmt.Println("watcher closed")
			return
		}
	}
}
