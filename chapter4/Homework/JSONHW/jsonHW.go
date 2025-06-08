package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type typicalStruct struct {
	Tel []int
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide a filename")
		return
	}

	filename := args[1]

	var myStruct typicalStruct
	err := readJSON(filename, &myStruct)
	if err != nil {
		fmt.Printf("error reading json: %v", err)
		return
	}

	for i := range myStruct.Tel {
		myStruct.Tel[i]++
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer f.Close()

	err = writeJSON(f, myStruct)
	if err != nil {
		fmt.Printf("error writing file: %v", err)
	}

	fmt.Println(myStruct)
}

func readJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	jsonDecoder := json.NewDecoder(in)
	err = jsonDecoder.Decode(key)
	if err != nil {
		return err
	}
	in.Close()

	return nil
}

func writeJSON(filename *os.File, key interface{}) error {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		return err
	}

	return nil
}
