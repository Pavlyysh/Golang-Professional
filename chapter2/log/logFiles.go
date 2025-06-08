package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "C:/Users/pgazukin/Desktop/go/Golang for Professional/log/logdir/log.txt"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
