package filedb

import (
	"encoding/json"

	"github.com/eriicafes/filedb/storage"
)

type ID int
type ForeignKey *ID

type Database struct {
	name    string
	storage storage.Storage
}

func New(name string) *Database {
	return &Database{
		name:    name,
		storage: storage.NewFileStorage(name),
	}
}

func (db *Database) Get(resource string, target any) {
	resourceData := db.storage.Get(resource)

	jsonData, err := json.Marshal(resourceData)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonData, target)
}

func (db *Database) Set(resource string, data []interface{}) {
	db.storage.Set(resource, data)
}
