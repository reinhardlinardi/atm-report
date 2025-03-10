package filestorage

type Storage interface {
	Get(path string) ([]byte, error)
}
