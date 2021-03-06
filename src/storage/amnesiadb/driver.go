/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package amnesiadb

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

func New() (Storage, error) {
	return Storage{}, nil
}

func (s Storage) Get(key string) (string, error) {
	if len(data) < 1 {
		return "", errors.New("There is currently no data.")
	}

	_, exists := data[key]
	if !exists {
		return "", errors.New("The key does not exist.")
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

func (s Storage) Remove(key string) error {
	_, err := s.Get(key)
	if err != nil {
		return err
	}

	delete(data, key)
	return nil
}
