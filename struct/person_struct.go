package main

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) getFullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.FirstName = parts[0]
	p.LastName = parts[1]
}

func main() {
	var person Person
	person = Person{"Fulano", "Tal"}

	fmt.Println(person.getFullName())

	person.setFullName("Algum Nome")
	fmt.Println(person.getFullName())
}
