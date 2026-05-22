package main

import "fmt"

func main() {

    var numero int

    fmt.Print("Digite um número: ")
    fmt.Scanln(&numero)

    if numero > 0 {
        fmt.Println("Número positivo")
		

    } else {
        fmt.Println("Número negativo")
    }

}