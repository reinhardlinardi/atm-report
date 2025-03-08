package fswatch

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

type WatcherImpl struct {
	watcher *fsnotify.Watcher
}

func New() (*WatcherImpl, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("err create watcher: %s\n", err.Error())
		return nil, err
	}

	w := &WatcherImpl{watcher: watcher}
	return w, nil
}

func (w *WatcherImpl) WatchCreated(ctx context.Context, path string, channel chan string) error {
	if err := w.watcher.Add(path); err != nil {
		fmt.Printf("err watch %s: %s\n", path, err.Error())
		return err
	}

	fmt.Println("watcher started")

	for {
		select {
		case event := <-w.watcher.Events:
			switch {
			// New file created
			case event.Op&fsnotify.Create == fsnotify.Create:
				filename := event.Name
				channel <- filename
			}
		case <-ctx.Done():
			w.watcher.Close()
			close(channel)

			fmt.Println("watcher stopped")
			return nil
		}
	}
}
