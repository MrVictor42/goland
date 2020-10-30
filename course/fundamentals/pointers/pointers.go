package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiro
	// não se pode fazer p++

	var p *int = nil // nulo

	p = &i //"pegue o endereço de i e me atribua esse endereço. Aqui P possui o endereço de memória de i "
	*p++   // desreferenciando (pegando o valor de p)
	i++

	//Como P e I compartilham o mesmo endereço, quando se faz *p++, se soma 1 em I e quando se faz o i++, está se somando 1 em P, por isso que a resposta é 3

	fmt.Println("Endereço de p:", p)
	fmt.Println("Valor de p:", *p)
	fmt.Println("Endereço de i:", &i)
	fmt.Println("Valor de i:", i)
}
