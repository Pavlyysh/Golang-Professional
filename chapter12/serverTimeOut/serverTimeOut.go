package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := ":8001"
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	} else {
		PORT = ":" + args[1]
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	}

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second, // время ожидания запроса клиента
		WriteTimeout: 3 * time.Second, // время на ответ клиенту (после конца чтения)
	}

	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}
