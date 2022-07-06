
# FileDB

A File DB for prototyping in Go


## Installation

To install `filedb` package, you need to install Go.

```bash
go get github.com/eriicafes/filedb
```

Import in your code:
```go
import "github.com/eriicafes/filedb"
```
    
## Usage/Examples

```go
package main

import (
	"github.com/eriicafes/filedb"
)

func main() {
	db := filedb.New()
    
	people := []string{"First", "Second", "Third"}

	// store to db
	var data []interface{}

	for _, person := range people {
		data = append(data, person)
	}

	db.Set("people", data)

	// retrieve from db
	var retrievedPeople []string

	db.Get("people", &retrievedPeople)
}
```

More advanced example with structs:

```go
package main

import (
	"fmt"

	"github.com/eriicafes/filedb"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db := filedb.New("db")

	users := []user{
		{Name: "Karen", Age: 23},
		{Name: "Andrew", Age: 12},
	}

	var data []interface{}

	for _, user := range users {
		data = append(data, user)
	}

	db.Set("users", data)

	var savedUsers []user

	db.Get("users", &savedUsers)

	fmt.Println(users)
}
```
