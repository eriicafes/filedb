package filedb

import (
	"encoding/json"
	"fmt"
	"os"
)

type Store map[string][]interface{}

type Storage interface {
	Get(resource string) (data []interface{})
	Set(resource string, data []interface{})
}

type fileStorage struct {
	name string
}

func NewFileStorage(name string) *fileStorage {
	return &fileStorage{name}
}

const fileStorageDefaultString = "{}"

func (fs *fileStorage) Filename() string {
	return fs.name + ".json"
}

func (fs *fileStorage) read() []byte {
	file, err := os.ReadFile(fs.Filename())

	if err != nil {
		fmt.Println("db read err:", err)

		file, err := os.Create(fs.Filename())

		file.WriteString(fileStorageDefaultString)

		file.Close()

		if err != nil {
			panic(err)
		}
		return fs.read()
	}

	return file
}

func (fs *fileStorage) write(data []byte) {
	err := os.WriteFile(fs.Filename(), data, os.FileMode(0644))

	if err != nil {
		fmt.Println("db write err:", err)

		panic(err)
	}
}

func (fs *fileStorage) reset() {
	fs.write([]byte(fileStorageDefaultString))
}

func (fs *fileStorage) Get(resource string) (data []interface{}) {
	// read data
	contents := fs.read()

	// bind to store
	var store Store

	if err := json.Unmarshal(contents, &store); err != nil {
		fs.reset()
		return fs.Get(resource)
	}

	return store[resource]
}

func (fs *fileStorage) Set(resource string, data []interface{}) {
	// read data
	contents := fs.read()

	// bind to store
	var store Store

	if err := json.Unmarshal(contents, &store); err != nil {
		fs.reset()
		fs.Set(resource, data)
		return
	}

	// update store
	store[resource] = data

	// store new data
	newContents, err := json.MarshalIndent(store, "", "   ")

	if err != nil {
		panic(err)
	}

	fs.write(newContents)
}
