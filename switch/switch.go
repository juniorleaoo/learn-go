package main

import ("fmt")

func main(){
    printType := func(i interface{}){
        switch i.(type){
            case bool:
                fmt.Println("bool");
            case int:
                fmt.Println("int");
            default: 
                fmt.Println("i don't");
        }
    }

    printType(true);
    printType(1);
    printType("haha");
}