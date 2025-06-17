// В последнем подразделе этого раздела показано, как можно организовать совмест-
// ный доступ к данным с помощью выделенной для этого горутины. Несмотря на
// то что общая память является традиционным способом взаимодействия потоков
// между собой, в стандартный комплект поставки Go входят встроенные функции
// синхронизации, которые позволяют одной горутине владеть общей частью дан-
// ных. Это означает, что другие горутины должны отправлять сообщения этой гору-
// тине, которая владеет общими данными, что предотвращает повреждение данных.
// Такая горутина называется управляющей горутиной. Согласно терминологии Go,
// это не коммуникация посредством общей памяти, а общая память посредством
// коммуникации.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup

	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}
