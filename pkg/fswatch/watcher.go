package fswatch

import (
	"context"

	"github.com/fsnotify/fsnotify"
)

type WatcherImpl struct {
	watcher *fsnotify.Watcher
}

func New() (*WatcherImpl, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	w := &WatcherImpl{watcher: watcher}
	return w, nil
}

func (w *WatcherImpl) WatchCreated(ctx context.Context, path string, channel chan string) error {
	if err := w.watcher.Add(path); err != nil {
		return err
	}

	for {
		select {
		case event := <-w.watcher.Events:
			switch {
			// New file created
			case event.Op&fsnotify.Create == fsnotify.Create:
				channel <- event.Name
			}
		case <-ctx.Done():
			w.watcher.Close()
			close(channel)

			return nil
		}
	}
}
