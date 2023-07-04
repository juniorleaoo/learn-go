package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.Price * (1 - p.Discount)
}

func main() {
	var product Product
	product = Product{
		Name:     "Teste",
		Price:    100,
		Discount: 0.05,
	}

	fmt.Println(product.priceWithDiscount())
}
