package aula_12_funcao_variatica

import (
	"fmt"
)

func soma(args ...int) int {
	total := 0
	for _, num := range args {
		total += num
	}
	return total
}

func FuncaoVariatica() {
	fmt.Println(soma(1, 2, 3))      // Saída: 6
	fmt.Println(soma(4, 5, 6, 7))   // Saída: 22
	fmt.Println(soma(10))           // Saída: 10
	fmt.Println(soma())             // Saída: 0
}