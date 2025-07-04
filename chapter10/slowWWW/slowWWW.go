package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	PORT := ":8001"
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Using default port number", PORT)
	} else {
		PORT = ":" + args[1]
	}

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}
