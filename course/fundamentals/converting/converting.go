package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Para conversão deve ser feita a tipagem (float64)
	*/
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ... O valor 123 é da tabela ASCII com o string antes
	fmt.Println("Teste " + string(123))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}
