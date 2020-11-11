package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //returno limpo. O retorno vai ser na ordem do código executado, sendo o primeiro o SEGUNDO e o segundo o PRIMEIRO. O retorno é dado na ordem dos parametros (segundo int, primeiro int)
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
