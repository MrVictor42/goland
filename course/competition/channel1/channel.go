package main

import "fmt"

/*
	O canal é um caminho de sincronismo. Aqui ele roda o caminho do canal, passa por ele enquanto executa as demais linhas. Mas tem que ter cuidado para não gerar deadlock
*/

func main() {
	ch := make(chan int, 1) //canal de inteiros

	ch <- 42 // enviando o numero 42 para o canal (escrita)
	<-ch     // recebendo dados do canal (leitura)
	ch <- 2
	fmt.Println(<-ch)
}
