/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package storage

type Storage interface {
	New(...string) Storage
	Get(string) (string, error)
	GetAll() ([]string, error)
	Put(string, string) error
}
