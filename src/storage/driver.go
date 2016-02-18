/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package storage

type Storage interface {
	New(...string) (Storage, error)
	Get(string) (string, error)
	GetAll() (map[string]string, error)
	Put(string, string) error
	Remove(string) error
}
