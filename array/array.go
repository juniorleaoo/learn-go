package main

import ("fmt")

var x [5]int

func main(){
	x[0] = 1
    x[1] = 10
	
	fmt.Println("-- X --")
    fmt.Println(x[0], x[1], x[2])
    fmt.Println(x)
	fmt.Println(len(x))
	
	y := [3]int{2, 4, 8}
	fmt.Println("-- Y --")
    fmt.Println(y)
	
	matriz := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("-- matriz --")
    fmt.Println(matriz)
}