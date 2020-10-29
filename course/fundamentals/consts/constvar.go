package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variavel. Forma mais utilizada. É preciso utilizar as variaveis,
	// declarar a variavel e não usar, da erro
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	e, f, g := 42, true, "bora"

	fmt.Println(e, f, g)
}
