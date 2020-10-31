package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      12455.42,
		"Gabriela Silva": 14252.42,
		"Pedro Junior":   1421.42,
	}

	//O primeiro elemento do map é a chave e depois vem o valor

	funcsESalarios["Victor Hugo"] = 20000.42

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
