package main

import "fmt"

func main() {

    var senha int

    for {

        fmt.Print("Digite a senha: ")
        fmt.Scan(&senha)

        if senha == 1234 {
            fmt.Println("Senha correta")
            break

        } else {
            fmt.Println("Senha incorreta")
        }

    }

}s