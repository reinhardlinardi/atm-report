package fswatch

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher
	path    string
}

func New() (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("err create watcher: %s\n", err.Error())
		return nil, err
	}

	w := &Watcher{watcher: watcher}
	return w, nil
}

func (w *Watcher) WatchCreated(ctx context.Context, path string, channel chan string) error {
	if err := w.watcher.Add(path); err != nil {
		fmt.Printf("err watch %s: %s\n", path, err.Error())
		return err
	}

	for {
		select {
		case event := <-w.watcher.Events:
			switch {
			// New file created
			case event.Op&fsnotify.Create == fsnotify.Create:
				filename := event.Name
				fmt.Println(filename)
				channel <- filename
			}
		case <-ctx.Done():
			w.watcher.Close()
			return nil
		}
	}
}
