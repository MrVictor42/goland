package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//Aqui se tem um slide referenciando 10 elementos de 20 posições
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
