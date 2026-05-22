package main

import "fmt"

func main() {

    numeros := [...]int{10, 45, 7, 89, 23}

    maior := numeros[0]

    for _, numero := range numeros {

        if numero > maior {
            maior = numero
        }

    }

    fmt.Println("Maior número:", maior)

}