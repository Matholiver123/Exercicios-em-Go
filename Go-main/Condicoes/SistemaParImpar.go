package main

import "fmt"

func main() {

    var numero int

    fmt.Print("Digite um número: ")
    fmt.Scanln(&numero)

	if numero % 2 == 0 {
        fmt.Println("Número par")
		

    } else {
        fmt.Println("Número impar")
    }

}