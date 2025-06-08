package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Paul",
		Surname: "Gazukin",
		Tel: []Telephone{Telephone{Mobile: true, Number: "123-12"},
			Telephone{Mobile: true, Number: "123-345"},
			Telephone{Mobile: true, Number: "abc-123"},
		},
	}

	saveToJSON(os.Stdout, myRecord)
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
