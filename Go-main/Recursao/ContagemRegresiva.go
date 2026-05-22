package main

import "fmt"

func contagem(numero int) {

    fmt.Println(numero)

    if numero > 0 {
        contagem(numero - 1)
    }

}

func main() {

    contagem(5)

}