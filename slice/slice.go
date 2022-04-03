package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "pêssego"}

	//Adiciona um elemento novo no slice
	slice = append(slice, "arroz")
	slice = append(slice, "feijão", "comida")

	fmt.Println("tamanho do slice:", len(slice))

	for indice, valor := range slice {
		fmt.Println("indice:", indice, "valor:", valor)
	}

	fmt.Println("- MAKE -")
	learnMake()
	fmt.Println("- COPY -")
	learnCopy()
	fmt.Println("- SLICE OPERATOR -")
	learnSliceOperator(slice)
	fmt.Println("- DELETE -")
	learnDeleteValueInSlice()
}

//deleta o 3 no slice
func learnDeleteValueInSlice() {
	sliceForDelete := []int{1, 2, 3, 4, 5}
	sliceDelete := append(sliceForDelete[:2], sliceForDelete[3:]...)
	fmt.Println(sliceDelete)
}

//operador slice, exibe os valores entre os índices 2 e 4
func learnSliceOperator(slice []string) {
	fmt.Println(slice[2:4])
	fmt.Println(slice[2:])
	fmt.Println(slice[:3])
}

//cria um slice vazio com tamanho inicial e valores non-zero
func learnMake() {
	sliceMake := make([]int, 3)
	sliceMake = append(sliceMake, 1, 2, 3)
	fmt.Println(sliceMake)
}

//faz uma cópia do slice
func learnCopy() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceCopy := make([]int, 6)
	copy(sliceCopy, slice)
	fmt.Println(sliceCopy)
}
