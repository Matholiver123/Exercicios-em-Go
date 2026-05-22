package main

import "fmt"

func crescente(numero int) {

    if numero > 5 {
        return
    }

    fmt.Println(numero)

    crescente(numero + 1)

}

func main() {

    crescente(1)

}