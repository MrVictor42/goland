package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	//O "_" é para ignorar o indice. Se tirar o _, você vai estar acesando o indice, e não o elemento
	for _, numero := range numeros { 
		fmt.Printf("%d ", numero)
	}
}