package main

import "fmt"

func somar(numero int) int {

    if numero == 1 {
        return 1
    }

    return numero + somar(numero-1)

}

func main() {

    resultado := somar(5)

    fmt.Println(resultado)

}