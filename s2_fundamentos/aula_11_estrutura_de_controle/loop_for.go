package aula_11_estrutura_de_controle

import (
	"fmt"
)

func LoopFor() {
	// Exemplo de for loop
	fmt.Println(">> Exemplos de for loop:")

	fmt.Println("Exemplo de for com única condição:")
	contador := 1
	for contador <= 5 {
		fmt.Println(contador)
		contador++
	}

	fmt.Println("Exemplo de for com inicialização, condição e pós-loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Exemplo de loop infinito com for:")
	contador2 := 0
	for {
		fmt.Println("Executando loop infinito...")
		contador2++
		if contador2 > 3 {
			break // Interrompe o loop após 3 iterações
		}
	}

}
