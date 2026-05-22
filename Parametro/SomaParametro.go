package main

import "fmt"

func somar(a int, b int) {

    resultado := a + b

    fmt.Printf("%v + %v = %v\n", a, b, resultado)

}

func main() {

    somar(10, 5)

}