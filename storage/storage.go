package storage

type Store map[string][]interface{}

type Storage interface {
	Get(resource string) (data []interface{})
	Set(resource string, data []interface{})
}
