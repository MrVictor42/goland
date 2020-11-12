package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n Misturando as estruturas de controle... ")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		//laço infinito
		fmt.Printf("\nPara sempre... a cada 5 segundos...")
		time.Sleep(time.Second * 5)
	}
}