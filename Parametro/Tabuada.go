package main

import "fmt"

func tabuada(numero int) {

    for i := 1; i <= 10; i++ {

        fmt.Printf("%v x %v = %v\n", numero, i, numero*i)

    }

}

func main() {

    tabuada(7)

}