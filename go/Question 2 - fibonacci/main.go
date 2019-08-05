package main

import "fmt"

func main() {
	Result := Solution(10)
	fmt.Println(Result)
}

func Solution(limit int) interface{} {
	var result [10]int
	for n := 0; n < limit; n++ {
		result[n] = fibonacci(n)
	}
	return result
}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		save := a
		a = b
		b = save + a
	}
	return a
}
