package main

import (
	"encoding/xml"
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
		fmt.Println("please give us a filename")
		return
	}

	filename := args[1]

	var myRecord Record
	err := loadFromXML(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("XML:", myRecord)
}

func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
	if err != nil {
		return err
	}
	in.Close()

	return nil
}
