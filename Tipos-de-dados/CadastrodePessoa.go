package main

import "fmt"

func main() {

    var nome string
    var idade int
    var altura float64
    var estudante bool

    fmt.Print("Digite seu nome: ")
    fmt.Scan(&nome)

    fmt.Print("Digite sua idade: ")
    fmt.Scan(&idade)

    fmt.Print("Digite sua altura: ")
    fmt.Scan(&altura)

    fmt.Print("É estudante? ")
    fmt.Scan(&estudante)

    fmt.Printf("\nNome: %v | Tipo: %T\n", nome, nome)
    fmt.Printf("Idade: %v | Tipo: %T\n", idade, idade)
    fmt.Printf("Altura: %v | Tipo: %T\n", altura, altura)
    fmt.Printf("Estudante: %v | Tipo: %T\n", estudante, estudante)

}