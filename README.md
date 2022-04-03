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

func main() {
	const number2 = 20
}

```

Obs.: Uma constante número não tem um tipo até receber um tipo explícito.

### Switch

É possível fazer comparações pelo tipo

```go
package main

import (
	"fmt"
)

func main() {
	printType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("i don't")
		}
	}

	printType(true)
	printType(1)
	printType("haha")
}
```

Obs.: Também é possível declarar funções dentro de funções

### Array

O número de elesmentos do array é fixo Existem algumas funções que podemos usar nos arrays, como o `len` para obter o
tamanho do array

```go
var x [2]int
len(x) //2
```

Arrays também podem ter mais de uma dimensão

```go
var two[2][2]int{{1, 2}, {3, 4}}
```

### Slice

A diferença do Slice para o Array NA declaração, é que array vai ter o seu tamanho finito específicado

```go
slice := []int{1, 2, 3, 4}
```

`append` adiciona um elemento no slice e retorna um novo slice com o novo elemento adicionado

`len` retorna o tamanho do slice

`range` fornece o índice e o valor de cada iteração do slice ou array

`make` cria um slice vazio com tamanho inicial e com valores non-zero. Make também cria um slice dinâmico, posso fazer
vários appends para ir adicionando novos valores no slice. O terceiro parâmetro do make é o `cap`, ele é sempre maior
que o `len`.

`copy` faz uma cópia da estrutura

`slice operator`

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice[2:4] //exibe os valores do slice entre os índices 2 e 3
```

Para deletar um valor no slice, é necessário recriar um novo slice 
