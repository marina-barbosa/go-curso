package aula_14_defer

import "fmt"

func conectar() {
	fmt.Println("Abrindo conexão com o banco de dados...")
}

func desconectar() {
	fmt.Println("Fecha conexão com o banco de dados.")
}

func manipularBanco() {
	fmt.Println("Manipulando banco: Executando consulta no banco de dados...")
	fmt.Println("Manipulando banco: Inserindo registros...")
}

func fimPrograma() {
	fmt.Println(" > Fim do programa.")
}

func Defer(deuErro bool) {

	fmt.Println(" >> Início do programa.")
	conectar()
	defer fimPrograma()
	defer desconectar()
	manipularBanco()

	if deuErro {
		fmt.Println("Deu algum erro.")
		return
	} else {
		fmt.Println("Operações concluídas.")
		return
	}
}
