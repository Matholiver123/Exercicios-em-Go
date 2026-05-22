
package main

import "fmt"

func Verificaridade(){
	var idade int
	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)
	if idade >= 18 {
		fmt.Println("Você é maior de idade.")	
	}else{
		fmt.Println("Você é menor de idade.")
	}
}

func main() {
Verificaridade()
  

}