package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) //o slice 2 possui duas posições
	copy(slice2, slice1) // copia os valores de slice 1 para slice 2. Obrigatóriamente deve ser um slice ou uma string
	fmt.Println(slice2)
}
