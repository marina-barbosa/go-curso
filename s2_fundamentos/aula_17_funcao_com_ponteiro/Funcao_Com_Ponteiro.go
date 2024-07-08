package aula_17_funcao_com_ponteiro

import "fmt"

func invertSign(num *int) {
	if num == nil {
		fmt.Println("Ponteiro nulo")
		return
	}
	*num = -*num
}

func FuncaoComPonteiro() {
	var n int = 10
	fmt.Println("Antes:", n) // Antes: 10
	invertSign(&n)
	fmt.Println("Depois:", n) // Depois: -10

	var m *int = nil
	invertSign(m) // Ponteiro nulo
}
