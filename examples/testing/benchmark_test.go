package testing

import (
	"errors"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	n := 35
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFastFib(b *testing.B) {
	n := 35
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastFib(n)
	}
}

func Fib(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n can not less than 1")
	}

	if n < 2 {
		return n, nil
	}
	i, _ := Fib(n - 1)
	j, _ := Fib(n - 2)
	return i + j, nil
}

func FastFib(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n can not less than 1")
	}
	cur, next := 0, 1
	for i := 0; i < n; i++ {
		cur, next = next, cur + next
	}
	return cur, nil
}