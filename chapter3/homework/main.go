package main

import (
	"errors"
	"fmt"
	"time"
)

type Weekday int

const (
	_ Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	// arr := [5]int{1, 1, 2, 3, 2}
	// mapArr := make(map[int]int)
	// for i, v := range arr {
	// 	mapArr[i] = v
	// }
	// fmt.Println(mapArr)

	// resDate := parseDate("2025-05-27")
	// fmt.Println(resDate)

	resTime := parseTime("00:29")
	fmt.Println(resTime)
}

func parseTime(s string) *time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		errors.New("incorrect date")
		return nil
	}
	return &t
}

func parseDate(s string) *time.Time {
	d, err := time.Parse("2006-01-02", s)
	if err != nil {
		errors.New("incorrect date")
		return nil
	}
	return &d
}
