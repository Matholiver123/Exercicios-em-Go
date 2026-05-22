package main

import "fmt"

type Pessoa struct  {
	nome string
	idade int
	altura float64
	peso float64
}

func main()  {
	pessoa := Pessoa{
		nome:  "Matheus",
        idade: 18,
		altura: 1.84,
		peso: 73.4,
}
   fmt.Println(pessoa.nome)
    fmt.Println(pessoa.idade)
	fmt.Println(pessoa.altura)
    fmt.Println(pessoa.peso)
}