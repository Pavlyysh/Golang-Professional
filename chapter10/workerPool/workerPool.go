// 		Пул обработчиков — это множество потоков, предназначенных для обработки на-
// значаемых им заданий. Веб-сервер Apache и Go-пакет net/http работают прибли-
// зительно так: основной процесс принимает все входящие запросы, которые затем
// перенаправляются рабочим процессам для обработки. Как только рабочий процесс
// завершает свою работу, он готов к обслуживанию нового клиента.
//
// 		Однако здесь есть главное различие: пул обработчиков использует не потоки,
// а горутины. Кроме того, потоки обычно не умирают после обработки запросов, по-
// тому что затраты на завершение потока и создание нового слишком высоки, тогда
// как горутина прекращает существовать после завершения работы. Как вы вскоре
// увидите, пулы обработчиков в Go реализованы с помощью буферизованных ка-
// налов, поскольку они позволяют ограничить число одновременно выполняемых
// горутин.

// 		Эта программа решает простую задачу: она обрабатывает целые числа и выводит их квадраты
// посредством отдельной горутины для обслуживания каждого запроса. Несмотря
// на простоту workerPool.go, код Go программы вполне можно использовать как
// шаблон для реализации более сложных задач.

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("need #jobs and #workers")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		// Не является необходимой, но дает лучшее представление
		// о том, как будут выводиться на экран результаты программы
		time.Sleep(time.Second)
	}
	wg.Done()
}

func makeWP(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}
