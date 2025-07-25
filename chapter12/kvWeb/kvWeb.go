package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type myElement struct {
	Name    string
	Surname string
	ID      string
}

var DATA = make(map[string]myElement)
var DATAFILE = "/tmp/dataFile.gob"

func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}

	PORT := ":8001"
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + args[1]
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("home.gohtml"))
	myT.ExecuteTemplate(w, "home.gohtml", nil)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")

	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">List all elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">Change an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">Insert new element!</a>")
	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")

	fmt.Fprintf(w, "<ul>")
	for k, v := range DATA {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of KV store!")
	tmpl := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !CHANGE(key, n) {
		fmt.Println("Update operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV store!")
	tmpl := template.Must(template.ParseFiles("insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")

	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !ADD(key, n) {
		fmt.Println("Add operation failed")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("cannot save to", DATAFILE)
		return err
	}

	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	if err != nil {
		fmt.Println("Error loading", DATAFILE)
		return err
	}
	defer loadFrom.Close()

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}

func ADD(key string, value myElement) bool {
	if key == "" {
		fmt.Println("invalid key")
		return false
	}
	if LOOKUP(key) == nil {
		DATA[key] = value
		return true
	}

	return false
}

func DELETE(key string) bool {
	if LOOKUP(key) != nil {
		delete(DATA, key)
		return true
	}

	return false
}

func CHANGE(key string, newValue myElement) bool {
	if LOOKUP(key) != nil {
		DATA[key] = newValue
		return true
	}

	return false
}

func LOOKUP(key string) *myElement {
	if key == "" {
		fmt.Println("invalid key")
		return nil
	}

	v, ok := DATA[key]
	if !ok {
		fmt.Println("key not found")
		return nil
	}
	return &v
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key: %s value: %v\n", k, v)
	}
}
