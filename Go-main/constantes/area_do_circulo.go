package main

import "fmt"

func main() {

    const PI float64 = 3.14

    var raio float64

    fmt.Print("Digite o valor do raio: ")
    fmt.Scan(&raio)

    area := PI * raio * raio

    fmt.Println("Área do círculo:", area)
}