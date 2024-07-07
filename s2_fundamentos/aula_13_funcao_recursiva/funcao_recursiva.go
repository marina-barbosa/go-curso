package aula_13_funcao_recursiva

import "fmt"

func soma(slice []int) int {
	// Condição de parada: slice vazio
	if len(slice) == 0 {
		return 0
	}

	// Caso recursivo: soma do primeiro elemento com a soma do restante
	return slice[0] + soma(slice[1:])
}

func fibonacci(pos int) int {
	// Condição de parada
	if pos <= 1 {
		return pos
	}

	// Chamada recursiva
	return fibonacci(pos-1) + fibonacci(pos-2)
}

func FuncaoRecursiva() {
	fmt.Println("- - função recursiva de soma - -")
	numeros := []int{1, 2, 3, 4, 5}
	resultado := soma(numeros)
	fmt.Printf("A soma dos elementos do slice é %d\n", resultado)
	// A soma dos elementos do slice é 15

	fmt.Println("- - funcao recursiva de fibonacci - -")
	termo := 15 // Calculando o décimo termo da sequência de Fibonacci
	resultadof := fibonacci(termo)
	fmt.Printf("O %dº termo da sequência de Fibonacci é %d\n", termo, resultadof)

	fmt.Println("- - printando a sequencia inteira - -")
	for i := int(1); i <= termo; i++ {
		fmt.Printf("O %dº termo da sequência de Fibonacci é %d\n", i, fibonacci(i))
	}
}
