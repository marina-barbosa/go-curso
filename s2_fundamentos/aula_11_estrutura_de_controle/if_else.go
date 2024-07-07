package aula_11_estrutura_de_controle

import (
	"fmt"
)

func IfElse() {

	// Exemplo de estrutura condicional if-else
	fmt.Println("\nExemplo de estrutura condicional if-else:")
	numero := 10
	if numero > 0 {
		fmt.Println("O número é positivo.")
	} else if numero < 0 {
		fmt.Println("O número é negativo.")
	} else {
		fmt.Println("O número é zero.")
	}

}