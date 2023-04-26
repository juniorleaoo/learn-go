package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //receiver memory address
	*p++
	fmt.Println(i)
	fmt.Println("memory address", p)
	fmt.Println("value", *p)
}
