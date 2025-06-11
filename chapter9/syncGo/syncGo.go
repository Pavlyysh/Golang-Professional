// Как дождаться завершения горутин,
// прежде чем закончить программу
// В этом разделе показано, как предотвратить завершение функции main(), пока
// она ожидает завершения своих горутин, с помощью пакета sync.

package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines\n", count)

	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Wait()
	fmt.Println("\nExiting...")
}
