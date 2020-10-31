package main

import "fmt"

func main() {
	//homogênea (mesmo tipo de dado) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	notasFinais := 0.0
	aux :=0

	for aux = 0; aux < len(notas); aux ++ {
		notasFinais += notas[aux]
	}

	notasFinais = notasFinais/ float64(len(notas))
	fmt.Printf("Media final é: %.2f\n", notasFinais)
}
