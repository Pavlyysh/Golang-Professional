package main

import (
	"testing"
	"testing/quick"
)

var N = 1_000_000

// Два вызова функции quick.Check() автоматически генерируют случайные числа
// на основе сигнатуры первого аргумента, который является функцией, определен-
// ной ранее. Создавать случайные входные числа самостоятельно не нужно, и это
// облегчает чтение и запись кода. Фактические тесты в обоих случаях выполняются
// в функции condition.
func TestWithSystem(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == (a + b)
	}
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestWithItself(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == Add(b, a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
