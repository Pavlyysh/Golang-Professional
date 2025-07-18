package main

import "fmt"

func main() {
	fmt.Println(fibo1(40))
	fmt.Println(fibo2(40))
	fmt.Println(fibo3(40))
}

// В этом коде содержится реализация функции fibo1(), которая использует
// рекурсию для вычисления чисел последовательности Фибоначчи. Алгоритм
// работает нормально, однако он довольно примитивный и несколько медленный.
func fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

// Здесь мы видим реализацию функции fibo2(), которая практически не отлича-
// ется от уже знакомой нам функции fibo1(). Однако интересно посмотреть, как не-
// значительное изменение кода — один оператор if вместо блока if else if — влияет
// на производительность функции.
func fibo2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo2(n-1) + fibo2(n-2)
}

// В представленной здесь функции fibo3() реализован совершенно иной подход
// с использованием отображения Go и цикла for. Неизвестно, действительно ли
// этот подход быстрее, чем две другие реализации.
func fibo3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}
