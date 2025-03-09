package storage

type Storage interface {
	Fetch(path string) ([]byte, error)
}
