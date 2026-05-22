package main

import "fmt"

func verificarNumero(numero int) {

    if numero%2 == 0 {

        fmt.Println("Número par")

    } else {

        fmt.Println("Número ímpar")

    }

}

func main() {

    verificarNumero(8)

}