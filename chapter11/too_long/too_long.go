// Как быть, если тестирование выполняется слишком долго
// или вообще не заканчивается
// На тот случай, если инструменту go test требуется слишком много времени для
// завершения работы или по какой-либо причине он вообще ее не заканчивает,
// существует параметр -timeout.

package main

import "time"

func sleepWithMe() {
	time.Sleep(5 * time.Second)
}

func getOne() int {
	return 1
}

func getTwo() int {
	return 2
}
