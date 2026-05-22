package main

import "fmt"

func fatorial(numero int) int {

    if numero == 1 {
        return 1
    }

    return numero * fatorial(numero-1)

}

func main() {

    resultado := fatorial(5)

    fmt.Println(resultado)

}