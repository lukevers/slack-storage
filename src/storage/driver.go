/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package storage

type Storage interface {
	New(...string) Storage
	Get(string) (error, string)
	GetAll() (error, []string)
	Put(string, string) error
}
