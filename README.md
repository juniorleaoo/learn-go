# learn-go

### Comandos

```sh
go run hello-world/hello-world.go
go build hello-world/hello-world.go
./hello-world
```

### Variáveis
Em GO, se usa `var` para declarar uma nova variável
`:=` é usado para declarar e inicializar uma variável

```go
var a = "a"

b := "b"
```

### Contantes
`const` é usado para declarar constantes

Em todos os lugares que `var` pode ser usado, o `const` também pode.

```go
package main

const number = 10
func main(){
    const number2 = 20
}

```
Obs.: Uma constante número não tem um tipo até receber um tipo explícito.

### switch
É possível fazer comparações pelo tipo 

```go
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
```

Obs.: Também é possível declarar funções dentro de funções