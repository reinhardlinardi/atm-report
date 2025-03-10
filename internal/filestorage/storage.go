package filestorage

import "os"

type StorageImpl struct{}

func New() *StorageImpl {
	return &StorageImpl{}
}

func (s *StorageImpl) Get(path string) ([]byte, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return raw, nil
}
