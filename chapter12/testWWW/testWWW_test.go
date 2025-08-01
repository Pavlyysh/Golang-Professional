package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("expected code %d, got %d", http.StatusOK, status)
	}

	expect := `Fine!`
	if rr.Body.String() != expect {
		t.Errorf("expected %s, got %s", expect, rr.Body.String())
	}
}

func TestStatusNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/StatusNotFound", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusNotFound)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("expected code %d, got %d", http.StatusNotFound, status)
	}
}
