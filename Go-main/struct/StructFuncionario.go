package main

import "fmt"

type Funcionario struct {
	Nome    string
	Idade   int
	Cargo   string
	Salario float64
}

func main() {

	funcionario := Funcionario{
		Nome:    "Matheus",
		Idade:   18,
		Cargo:   "Programador Go",
		Salario: 3500.00,
	}

	fmt.Println("Nome:", funcionario.Nome)
	fmt.Println("Idade:", funcionario.Idade)
	fmt.Println("Cargo:", funcionario.Cargo)
	fmt.Println("Salário:", funcionario.Salario)

}