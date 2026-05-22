package main

import "fmt"

func potencia(base int, expoente int) int {

    if expoente == 0 {
        return 1
    }

    return base * potencia(base, expoente-1)

}

func main() {

    resultado := potencia(2, 4)

    fmt.Println(resultado)

}