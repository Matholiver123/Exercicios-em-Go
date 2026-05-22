package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número para ver a tabuada: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)

}
}