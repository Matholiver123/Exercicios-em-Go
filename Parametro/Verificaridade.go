package main

import "fmt"

func verificarIdade(idade int) {

    if idade >= 18 {

        fmt.Println("Maior de idade")

    } else {

        fmt.Println("Menor de idade")

    }

}

func main() {

    verificarIdade(20)

}