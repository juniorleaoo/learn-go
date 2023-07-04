package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"nome"`
	Price float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := Product{1, "Notebook", 1899.90, []string{"Test", "hi"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 Product
	p2StringJson := `{"id": 2,"nome":"Nome 2"}`
	json.Unmarshal([]byte(p2StringJson), &p2)
	fmt.Println(p2.ID)
	fmt.Println(p2.Name)
	fmt.Println(p2.Tags)
}
