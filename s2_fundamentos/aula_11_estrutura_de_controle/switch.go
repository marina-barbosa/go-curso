package aula_11_estrutura_de_controle

import (
	"fmt"
)

func Switch() {

	// Exemplo de switch
	fmt.Println("\nExemplo de switch:")
	dia := "quarta"
	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("É um dia útil.")
	case "sábado", "domingo":
		fmt.Println("É final de semana.")
	default:
		fmt.Println("Dia inválido.")
	}
}
