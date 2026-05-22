package main

import "fmt"

func main() {

    numeros := [...]int{10, 20, 30, 40, 50}

    for i := 0; i < len(numeros); i++ {

        fmt.Printf("Posição %v = %v\n", i, numeros[i])

    }

}