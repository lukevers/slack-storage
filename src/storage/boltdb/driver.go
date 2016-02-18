/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package boltdb

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"storage"
)

var BUCKET = []byte("slack-storage")

type Storage struct {
	storage.Storage
	db *bolt.DB
}

func New(dsn string) (Storage, error) {
	db, err := bolt.Open(dsn, 0600, nil)
	if err != nil {
		return Storage{}, fmt.Errorf("Error opening boltdb database: %s", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BUCKET)
		if err != nil {
			return fmt.Errorf("Error creating boltdb bucket: %s", err)
		}

		return nil
	})

	if err != nil {
		return Storage{}, err
	}

	return Storage{
		db: db,
	}, nil
}

func (s Storage) Get(key string) (string, error) {
	var value string
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(BUCKET)
		if key, _ := b.Cursor().First(); key == nil {
			return errors.New("There is currently no data.")
		}

		v := b.Get([]byte(key))
		if v == nil {
			return errors.New("The key does not exist.")
		}

		value = string(v)

		return nil
	})

	return value, err
}

func (s Storage) GetAll() (map[string]string, error) {
	data := make(map[string]string)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(BUCKET)
		b.ForEach(func(key, value []byte) error {
			data[string(key)] = string(value)
			return nil
		})

		if len(data) < 1 {
			return errors.New("There is currently no data.")
		}

		return nil
	})

	return data, err
}

func (s Storage) Put(key, value string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(BUCKET)
		return b.Put([]byte(key), []byte(value))
	})
}

func (s Storage) Remove(key string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(BUCKET)
		return b.Delete([]byte(key))
	})
	return err
}
