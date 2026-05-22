package main

import "fmt"

func main() {

    var nome string = "Matheus"
    var idade int = 18
    var altura float64 = 1.82
    var estudante bool = true

    fmt.Printf("Nome: %v\n", nome)
    fmt.Printf("Tipo: %T\n", nome)

    fmt.Printf("Idade: %v\n", idade)
    fmt.Printf("Tipo: %T\n", idade)

    fmt.Printf("Altura: %v\n", altura)
    fmt.Printf("Tipo: %T\n", altura)

    fmt.Printf("Estudante: %v\n", estudante)
    fmt.Printf("Tipo: %T\n", estudante)

}