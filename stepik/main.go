package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y string
	fmt.Scan(&x, &y)

	parsedX, _ := time.Parse("2006-01-02", x)
	parsedY, _ := time.Parse("2006-01-02", y)
	diff := parsedX.Year() - parsedY.Year()
	fmt.Printf("%d year ago", diff)
}
