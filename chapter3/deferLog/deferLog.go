package main

import (
	"log"
	"os"
)

const LOGFILE = "log.txt"

func one(aLog *log.Logger) {
	aLog.Println("----FUNC ONE---")
	defer aLog.Println("----FUNC ONE---")

	for i := 1; i <= 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("----FUNC TWO")
	defer aLog.Println("FUNC TWO---")

	for i := 1; i <= 10; i++ {
		aLog.Println(i)
	}
}

func main() {
	file, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	iLog := log.New(file, "iLog ", log.LstdFlags)

	iLog.Println("iLog HELLO")
	iLog.Println("Another HELLO")

	one(iLog)
	two(iLog)
}
