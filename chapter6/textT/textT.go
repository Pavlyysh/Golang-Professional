package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a template file")
		return
	}

	tFile := arguments[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []Entry

	for _, v := range DATA {
		if len(v) == 2 {
			temp := Entry{Number: v[0], Square: v[1]}
			Entries = append(Entries, temp)
		}
	}
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
