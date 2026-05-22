package main

import "fmt"

func main() {

    var nota int

    fmt.Print("Digite um número: ")
    fmt.Scanln(&nota)

	if nota >= 7 {
        fmt.Println("Aprovado")

		if nota >= 7 {

			} else if nota >= 5 {
				fmt.Println("Recuperação")
			
			}else {
        fmt.Println("Reprovado")
    }

}
}