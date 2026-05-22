package main

import "fmt"

func main() {

	idades := map[string]int{
		"Matheus": 18,
		"Gabriel": 20,
		"Lucas":   22,
	}

	fmt.Println("Idade do Matheus:", idades["Matheus"])
	fmt.Println("Idade do Gabriel:", idades["Gabriel"])
	fmt.Println("Idade do Lucas:", idades["Lucas"])

}