package fswatch

import "context"

type Watcher interface {
	WatchCreated(ctx context.Context, path string, channel chan string) error
}
