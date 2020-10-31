package main

import "fmt"

func main() {
	//var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12098765432] = "Draven"
	aprovados[23452129890] = "Joel"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12098765432])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados)
}
