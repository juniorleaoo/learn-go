package main

import (
	"fmt"
)

type User struct {
	Age  int
	Name string
}

func main() {
	users := make(map[int]string)

	users[11] = "Fulano"
	fmt.Println(users)
	fmt.Println(users[11])

	usersAsKey := make(map[User]string)
	usersAsKey[User{Age: 1, Name: "Fulano"}] = "Teste"
	usersAsKey[User{Age: 2, Name: "Fulano 2"}] = "Teste 2"
	fmt.Println("usersAsKey: ", usersAsKey)
	usersAsKey[User{Age: 2, Name: "Fulano 2"}] = "Teste 3"
	fmt.Println("usersAsKey: ", usersAsKey, len(usersAsKey)) // result: 2

	usersAsValue := make(map[string]User)
	usersAsValue["teste"] = User{Age: 1, Name: "Fulano"}
	usersAsValue["teste 1"] = User{Age: 2, Name: "Fulano 2"}
	fmt.Println("usersAsValue: ", usersAsValue)

	for name, user := range usersAsValue {
		fmt.Println(name, user)
	}

	usersAligned := map[string]map[string]User{
		"A": {
			"Fulano": User{Age: 1, Name: "Maria"},
		},
	}
	fmt.Println(usersAligned)

	for tipo, name := range usersAligned {
		for names, user := range name {
			fmt.Println(tipo, name, names, user)
		}
	}
}
