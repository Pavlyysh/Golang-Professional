// В этом примере мы cоздадим HTTP-клиент, который не желает слишком долго ждать ответа от HTTP-
// сервера, — вполне обычный сценарий. Фактически, поскольку почти все HTTP-
// клиенты поддерживают такую функциональность, мы рассмотрим другой способ
// отмены HTTP-запроса в главе 12.
//
// 		Программа useContext.go принимает два аргумента командной строки: URL-
// адрес сервера, к которому она будет подключаться, и время, в течение которого
// представленная утилита должна ожидать отклика сервера. Если передать про-
// грамме только один аргумент командной строки, то время ожидания составит пять
// секунд.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myURL string
	delay int = 5
	wg    sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}
	myURL = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)

	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s \n", myURL)
	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}

func connect(c context.Context) error {
	defer wg.Done()
	data := make(chan myData, 1)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", myURL, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled!")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", realHTTPData)
	}
	return nil
}
