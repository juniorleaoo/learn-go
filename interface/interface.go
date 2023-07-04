package main

import "fmt"

type Print interface {
	toString() string
}

type Person struct {
	Name string
}

type Product struct {
	Name string
}

func (p Person) toString() string {
	return p.Name + " Person"
}

func (p Product) toString() string {
	return p.Name + " Product"
}

func printInteface(p Print) string {
	return p.toString()
}

func main() {
	fmt.Println(printInteface(Person{"Fulano"}))
	fmt.Println(printInteface(Product{"Produto"}))
}
