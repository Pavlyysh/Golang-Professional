package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	database := arguments[1]
	tFile = arguments[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Emptying database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?,?,?)")
	for i := 2; i < 50; i++ {
		_, _ = stmt.Exec(i, i*2, i*i)
	}

	rows, err := db.Query("SELECT * FROM")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var n, d, s int

	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		if err != nil {
			fmt.Println(nil)
			return
		}
		temp := Entry{Number: n, Double: d, Square: s}
		DATA = append(DATA, temp)
	}

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}
