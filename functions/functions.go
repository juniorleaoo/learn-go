package main

import "fmt"

func main() {
	first, second := increment()
	fmt.Printf("first: %d   second: %d\n", first, second)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Printf("sub: %d", calc(2, 1, sub)) //1
}

func increment() (first int, second int) {
	first = 1
	second = 2
	return
}

func calc(a, b int, operator func(a, b int) int) int {
	return operator(a, b)
}
