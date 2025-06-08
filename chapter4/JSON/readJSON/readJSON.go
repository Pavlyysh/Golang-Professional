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
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no filename")
		os.Exit(1)
	}

	filename := args[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myRecord)
	}
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}
