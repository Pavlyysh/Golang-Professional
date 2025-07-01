// Тестирование программного обеспечения может
// показать только наличие одной или нескольких ошибок, но не отсут-
// ствие ошибок. Это означает, что вы никогда не можете быть абсолютно
// уверены, что в вашем коде нет ошибок!

package testMe

func f1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2 // специально допущена ошибка
	}
	return f2(n-1) + f2(n-2)
}

func s1(s string) int {
	if s == "" {
		return 0
	}
	n := 1 // специально допущена ошибка
	for range s {
		n++
	}
	return n
}

func s2(s string) int {
	return len(s)
}
