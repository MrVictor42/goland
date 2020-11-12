package main

import "fmt"

// revisão: um ponteiro é representado por um *

func inc1(n int) {
	n++ //n = n + 1
}

// revisão: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
func inc2(n *int) {
	*n++ // aqui pega o valor de n e soma mais um. Como tem o *, é possivel somar mais um
}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) //por referência
	fmt.Println(n)
}
