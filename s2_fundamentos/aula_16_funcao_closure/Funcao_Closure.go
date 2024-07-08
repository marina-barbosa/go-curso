package aula_16_funcao_closure

import "fmt"

// calcula impar
func calculateOdd() func() int {
	num := 1
	// returns inner function
	return func() int {
		num = num + 2
		return num
	}
}

// update status to-do
func status() func() string {
	var index int
	orderStatus := map[int]string{
		1: "TO DO",
		2: "DOING",
		3: "DONE",
	}
	// Retorna função que tem acesso à todo o escopo da função status, portanto também pode atualizar o valor da variável index
	return func() string {
		index++
		return orderStatus[index]
	}
}

// contador
func createCounter() func() int {
	count := 0
	incremento := func() int {
		count++
		return count
	}
	return incremento
}

// closure com parametro
func adder(x int) func(int) int {
	return func(y int) int {
			return x + y
	}
}

func FuncaoClosure() {

	odd := calculateOdd()  // call the outer function
	fmt.Println(odd())     // 3
	fmt.Println(odd())     // 5
	fmt.Println(odd())     // 7
	odd2 := calculateOdd() // call the outer function again
	fmt.Println(odd2())    // 3
	///////////////////
	update1 := status()
	update2 := status()
	fmt.Println(update1()) // to do
	fmt.Println(update1()) // doing
	fmt.Println(update1()) // done
	fmt.Println(update2()) // to do
	fmt.Println(update2()) // doing
	/////////////////////
	counter1 := createCounter()
	counter2 := createCounter()
	fmt.Println(counter1()) // Saída: 1
	fmt.Println(counter1()) // Saída: 2
	fmt.Println(counter2()) // Saída: 1
	fmt.Println(counter2()) // Saída: 2
	//////////////////////////////
	addFive := adder(5) // x = 5
	fmt.Println(addFive(3)) // y = 3, resultado: 8

	addTen := adder(10) // x = 10
	fmt.Println(addTen(7)) // y = 7, resultado: 17
}
