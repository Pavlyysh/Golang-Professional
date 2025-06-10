// В этом подразделе показано, как обрабатывать два сигнала в программе Go
// Программа handleTwo.go обрабатывает сигналы SIGINFO и SIGINT,
// которые в Go называются syscall.SIGINFO и os.Interrupt соответственно.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINT:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(time.Second * 20)
	}
}

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}
