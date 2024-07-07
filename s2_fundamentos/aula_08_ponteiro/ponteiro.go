package aula_08_ponteiro

import "fmt"

func Ponteiro() {
	// ponteiro é uma referencia de memória
	var varA int = 10
	var varB int = varA
	fmt.Println(varA, varB) // 10 10

	varA += 33
	println(varA, varB) // 43 10

	var varC int = 20
	var ponteiro *int = &varC

	println(varC, *ponteiro) // 20 20
	varC += 50
	println(varC, *ponteiro) // 70 70

	fmt.Printf("Valor de 'varC': %d\n", varC)                    // Valor de 'varC': 70
	fmt.Printf("Endereço de 'varC': %x\n", &varC)                // Endereço de 'varC': endereço de memória
	fmt.Printf("Valor de 'ponteiro': %x\n", ponteiro)            // Valor de 'ponteiro' : endereço de memória
	fmt.Printf("Valor apontado por 'ponteiro': %d\n", *ponteiro) // Valor apontado por 'ponteiro': 70

	*ponteiro = 20 // altera o valor de 'varC' através do 'ponteiro'

	fmt.Printf("Novo valor de 'varC': %d\n", varC) // Novo valor de 'varC': 20
}
