package aula_09_array_and_slice

import "fmt"

func ArrayAndSlice() {

	// Definir um array
	arr := [7]int{10, 20, 30, 40, 50, 60, 70}

	// Criar um slice a partir do array
	slice1 := arr[2:5] //O primeiro indice (2) é inclusivo, ou seja, ela vai pegar e o segundo indice (5) é exclusivo, ou seja, ele vai parar antes e não vai pegar.

	// Exibir o array e o slice
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice1) // [30, 40, 50]

	// Criar um slice
	slice2 := make([]int, 3) // [0 0 0]

	// Adicionar elementos
	slice2 = append(slice2, 4) // [0 0 0 4]

	// Copiar slices
	newSlice := make([]int, len(slice2))
	copy(newSlice, slice2) // newSlice = [0 0 0 4]

	// Subslice
	subslice := slice2[2:4] // [0 4]  // O primeiro indice (2) é inclusivo, ou seja, ela vai pegar e o segundo indice (4) é exclusivo, ou seja, ele vai parar antes e não vai pegar.

	// Comprimento e capacidade
	fmt.Println(len(slice2)) // 4
	fmt.Println(cap(slice2)) // 6 (dependente da alocação interna)
	fmt.Println(subslice)    // [0 4]
}
