package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var soma int

	fmt.Println("Digite um número:")
	fmt.Scanln(&num1)
	fmt.Println("Digite outro número:")
	fmt.Scanln(&num2)
	soma = num1 + num2

	fmt.Println("A soma dos números será:", soma)
}
