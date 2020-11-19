package main

import (
	"fmt"
	"time"
)

/*
	Channel (canal) - é a forma de comunicação entre goroutines
	É um tipo
	O channel só passa para as linhas subsequentes, quando os valores internos dele são preenchidos, então a Linha 28 é lida, é preenchido os valores para A e B, depois é printado B com os restante dos valores da função doisTresQuatroVezes 
*/

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second * 3)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println(<-c)
}
