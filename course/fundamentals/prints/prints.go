package main

import "fmt"

func main() {

	//O comando fmt.Print imprime a string na mesma linha
	fmt.Print("Mesma ")
	fmt.Print("linha.")

	//O comando fmt.Println imprime a string e no final quebra a linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	fmt.Println("O valor de x é:", x)

	//a função "Sprint(), pega um ponto flutuante e transforma em uma string"
	xs := fmt.Sprint(x)
	fmt.Println("o valor de x é: " + xs)

	//O printf permite formatar os valores, diferente do print e println
	fmt.Printf("o valor de x é: %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"

	//Para printar o valor de um boolean e um printf, se usa o %t
	fmt.Printf("\n%d %f %.1f %t %s\n", a, b, b, c ,d)
	fmt.Println(!c)
}