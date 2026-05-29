package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome       string  `json:"nome"`
	Idade      int     `json:"idade"`
	Altura     float64 `json:"altura"`
	Estudando  bool    `json:"estudando"`
	Cidade     string  `json:"cidade"`
}

func main() {

	var pessoa Pessoa

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&pessoa.Nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&pessoa.Idade)

	fmt.Print("Digite sua altura: ")
	fmt.Scanln(&pessoa.Altura)

	fmt.Print("Você está estudando? (true/false): ")
	fmt.Scanln(&pessoa.Estudando)

	fmt.Print("Digite sua cidade: ")
	fmt.Scanln(&pessoa.Cidade)

	dados, err := json.MarshalIndent(pessoa, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON")
		return
	}

	err = os.WriteFile("pessoa.json", dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo")
		return
	}

	fmt.Println("Pessoa salva com sucesso!")
}