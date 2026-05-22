package main

import "fmt"

func main() {

    numeros := [...]int{10, 20, 30, 40, 50}

    soma := 0

    for _, numero := range numeros {

        soma += numero

    }

    fmt.Println(soma)

}