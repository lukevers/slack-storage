/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package memory

import (
	"storage"
)

type Storage struct {
	storage.Storage
}

func New() Storage {
	return Storage{}
}

func (s Storage) Get(key string) (string, error) {
	return "x", nil
}

func (s Storage) GetAll() ([]string, error) {
	return nil, nil
}

func (s Storage) Put(key, value string) error {
	return nil
}
