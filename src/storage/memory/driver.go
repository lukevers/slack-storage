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

func (s Storage) Get(key string) (error, string) {
	return nil, "x"
}

func (s Storage) GetAll() (error, []string) {
	return nil, nil
}

func (s Storage) Put(key, value string) error {
	return nil
}
