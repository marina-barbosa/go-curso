package aula_09_array_and_slice

import "fmt"

func CapacidadeSlice() {
	// Criar um slice inicial vazio com capacidade 3
	slice := make([]int, 0, 3)

	// Função para exibir o comprimento e capacidade do slice
	printSliceInfo := func(s []int) {
		fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", s, len(s), cap(s))
	}

	// Adicionar elementos ao slice e observar a capacidade
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		printSliceInfo(slice)
	}
}
