// В этом подразделе показано, как обрабатывать все сигналы, но реагировать только
// на те, которые действительно вас интересуют. Это гораздо лучший и более без-
// опасный метод, чем тот, что был описан в handlerTwo.go.

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
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handle(sig)
			case syscall.SIGTERM:
				handle(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall.SIGUSR2!")
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(30 * time.Second)
	}
}

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}
