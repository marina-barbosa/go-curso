package aula_15_panic_and_recover

import "fmt"

func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	defer func() {
		fmt.Println("Panic chama todos os Defer antes.")
	}()
	erroCritico := true
	fmt.Println("Start")
	if erroCritico {
		panic("Something went wrong!")
	}
	fmt.Println("End") // Não será executado
}
