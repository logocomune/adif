# Adif Parser

Simple parser for Amateur Data Interchange Format (ADIF)

## Installation

`go get -u github.com/logocomune/adif`

## Usage

### Read from file

```go 
package main

import (
	"github.com/logocomune/adif"
	"fmt"
	"log"
)

func main() {

	parsed, err := adif.ParseFile("example.adi")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", parsed)
}


```

### Read from a string

```go
package main

import (
	"github.com/logocomune/adif"
	"fmt"
	"log"
)

const adifString = "...AN ADIF STRING..."

func main() {

	parsed, err := adif.ParseString(adifString)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Printf("%+v\n", parsed)
}
```
