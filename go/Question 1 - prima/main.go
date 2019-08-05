package main

import "fmt"

func main() {
	result := Solution(10)
	fmt.Println(result)
}

func Solution(limit int) map[int]int {
	var result = make(map[int]int)
	index := 0
	for i := 1; i <= limit; i++ {
		a := 0
		for y := 1; y <= i; y++ {
			if i%y == 0 {
				a++
			}
		}
		if a == 2 {
			result[index] = i
			index++
		}
	}
	return result
}
