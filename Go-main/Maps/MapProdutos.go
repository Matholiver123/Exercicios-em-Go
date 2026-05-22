package main

import "fmt"

func main() {

	produtos := map[string]float64{
		"Mouse":    89.90,
		"Teclado":  120.50,
		"Monitor":  950.00,
	}

	fmt.Println("Produto: Mouse")
	fmt.Println("Preço:", produtos["Mouse"])

	fmt.Println()

	fmt.Println("Produto: Teclado")
	fmt.Println("Preço:", produtos["Teclado"])

	fmt.Println()

	fmt.Println("Produto: Monitor")
	fmt.Println("Preço:", produtos["Monitor"])

}