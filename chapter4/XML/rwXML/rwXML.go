// Cчитывает с диска запись JSON, вносит в нее изменения,
// преобразует ее в формат XML и выводит на экран. Затем данные
// XML преобразуются в формат JSON.

package main

import (
	"encoding/json"
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
		fmt.Println("provide a filename please")
		return
	}

	filename := args[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JSON:", myRecord)

	myRecord.Name = "Pasha"

	xmlData, _ := xml.MarshalIndent(myRecord, "", "	")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Print("\nxmlData:", string(xmlData))

	data := &Record{}

	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println("error unmarshalling xml", err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshalling json", err)
		return
	}

	_ = json.Unmarshal(result, &myRecord)
	fmt.Println("\nJSON:", myRecord)
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
