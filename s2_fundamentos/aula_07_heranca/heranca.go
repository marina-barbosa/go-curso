package aula_07_heranca

import "fmt"

type Endereco struct {
	Rua    string
	Cidade string
}

// Embutindo o struct Endereco em Pessoa
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco // Composição
}

func Heranca() {
	// Criando uma instância de Pessoa
	p := Pessoa{
		Nome:  "Isaac",
		Idade: 32,
		Endereco: Endereco{
			Rua:    "123 Rua Principal",
			Cidade: "Santos",
		},
	}

	// Acessando os campos do struct embutido diretamente
	fmt.Println(p.Nome)   // Isaac
	fmt.Println(p.Idade)  // 32
	fmt.Println(p.Rua)    // 123 Rua Principal
	fmt.Println(p.Cidade) // Santos
}
