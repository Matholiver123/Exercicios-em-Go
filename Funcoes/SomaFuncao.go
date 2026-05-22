package main

import "fmt"

func somarFuncao(){
	var num1, num2 int
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)
	soma := num1 + num2
	fmt.Println("A soma dos número é: ", num1, num2, "=", soma)
}

func main() {
somarFuncao()
  

}