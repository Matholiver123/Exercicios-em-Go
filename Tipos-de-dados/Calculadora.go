package main

import "fmt"

func main() {

	var num1 int 
	var num2 int

	fmt.Print("Digite um número: ")
    fmt.Scan(&num1)

	fmt.Print("Digite outro número:")
	fmt.Scan(&num2)

	soma := num1 + num2
    subtracao := num1 - num2
    multiplicacao := num1 * num2
    divisao := num1 / num2

    fmt.Printf("A soma de %v e %v é: %v\n", num1, num2, soma)

    fmt.Printf("A subtração de %v e %v é: %v\n", num1, num2, subtracao)

    fmt.Printf("A multiplicação de %v e %v é: %v\n", num1, num2, multiplicacao)

    fmt.Printf("A divisão de %v e %v é: %v\n", num1, num2, divisao)



}