package main

import (
	"fmt"
)

func main() {
	var idade int

	fmt.Println("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade!")
	}

}
