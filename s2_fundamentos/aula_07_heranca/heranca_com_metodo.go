package aula_07_heranca

import "fmt"

// Struct básica
type Animal struct {
	Nome string
}

// Método para Animal
func (a Animal) Falar() {
	fmt.Println("Som de animal")
}

// Outra struct que embute Animal
type Cachorro struct {
	Animal
	Raca string
}

// Método específico para Cachorro
func (c Cachorro) Falar() {
	fmt.Println("Au Au! Meu nome é", c.Nome)
}

func HerancaComMetodo() {
	c := Cachorro{
		Animal: Animal{Nome: "Rex"},
		Raca:   "Golden Retriever",
	}
	c.Falar()        // Au Au! Meu nome é Rex
	c.Animal.Falar() // Som de animal
}
