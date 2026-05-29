package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Livro struct {
	Nome  string  `json:"nome"`
	Autor string  `json:"autor"`
	Preco float64 `json:"preco"`
}

func main() {

	livro := Livro{
		Nome:  "Diario",
		Autor: "Pedro",
		Preco: 29.99,
	}

	dados, err := json.MarshalIndent(livro, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON")
		return
	}

	err = os.WriteFile("livro.json", dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo")
		return
	}

	fmt.Println("Livro salvo com sucesso!")
}