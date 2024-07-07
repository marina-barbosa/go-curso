package aula_10_map

import "fmt"

func Map() {
	idade := map[string]int{"Isaac": 30, "Hortencia": 25}

	idade["Jane"] = 40

	for chave, valor := range idade {
		fmt.Printf("%s tem %d anos\n", chave, valor)
	}

	delete(idade, "Jane")

	valor, ok := idade["Jane"]
	if ok {
		fmt.Println("Idade de Jane:", valor)
	} else {
		fmt.Println("Jane não encontrada")
	}
}
