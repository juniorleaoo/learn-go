package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice))

	doubleslice := make([]int, 0, 2)
	fmt.Println(doubleslice, len(doubleslice), cap(doubleslice))

	doubleslice = append(doubleslice, 1, 2)
	fmt.Println(doubleslice, len(doubleslice), cap(doubleslice))

	doubleslice = append(doubleslice, 3)
	fmt.Println(doubleslice, len(doubleslice), cap(doubleslice))
}
