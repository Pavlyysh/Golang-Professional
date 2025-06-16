// 		Атомарная операция — это операция, которая выполняется за один шаг относи-
// тельно других потоков или в данном случае других горутин. Таким образом, ато-
// марная операция не может быть прервана на середине.

// 		В стандартную библиотеку Go входит пакет atomic, который иногда позволяет
// обойтись без использования мьютекса. Однако мьютексы более универсальны,
// чем атомарные операции. С помощью пакета atomic можно создавать атомные
// счетчики, доступные для нескольких горутин, не рискуя получить проблемы син-
// хронизации и состояние гонки.

// 		Обратите внимание, что в случае атомарной переменной все операции чтения
// и записи должны выполняться с применением атомарных функций, предоставля-
// емых пакетом atomic.

package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	val int64
}

func main() {
	minusX := flag.Int("x", 100, "goroutines")
	minusY := flag.Int("y", 200, "value")
	flag.Parse()
	X := *minusX
	Y := *minusY

	var wg sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < X; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Value())
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
