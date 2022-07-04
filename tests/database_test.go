package tests

import (
	"testing"

	"github.com/eriicafes/filedb"
)

func TestDatabaseGetSet(t *testing.T) {
	db := filedb.New("store")

	people := []string{"First", "Second", "Third"}

	peopleResource := "people"

	// store to db
	var data []interface{}

	for _, person := range people {
		data = append(data, person)
	}

	db.Set(peopleResource, data)

	// retrieve from db
	var retrievedPeople []string

	db.Get(peopleResource, &retrievedPeople)

	// check if data is same length
	if len(retrievedPeople) != len(people) {
		t.Error("retrieved data is not same length with original data")
	}
}
