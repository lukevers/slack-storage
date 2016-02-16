/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package memory

import (
	"errors"
	"storage"
)

var data map[string]string

func init() {
	data = make(map[string]string)
}

type Storage struct {
	storage.Storage
}

func New() Storage {
	return Storage{}
}

func (s Storage) Get(key string) (string, error) {
	if len(data) < 1 {
		return "", errors.New("There is currently no data.")
	}

	return "", nil
}

func (s Storage) GetAll() (map[string]string, error) {
	if len(data) < 1 {
		return nil, errors.New("There is currently no data.")
	}

	return data, nil
}

func (s Storage) Put(key, value string) error {
	data[key] = value
	return nil
}
