package main

import "testing"

func TestGetOne(t *testing.T) {
	sleepWithMe()
	value := getOne()
	if value != 1 {
		t.Errorf("func returned %v", value)
	}
	sleepWithMe()
}

func TestGetTwo(t *testing.T) {
	sleepWithMe()
	value := getTwo()
	if value != 2 {
		t.Errorf("func returned %v", value)
	}
}

func TestWithFail(t *testing.T) {
	value := getOne()
	if value != 2 {
		t.Errorf("func returned %v", value)
	}
}
