package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice -> tamanho variavel

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5] int {1, 2, 3, 4, 5}

	//Slice não é um array! Slide define um pedaço de um array. Um array pode ter um ou mais slices!
	s2 := a2[1:3] //Pega o indice 1 (2) e vai até o indice 3 (4), mas não incluindo o indice 3, então até o valor 3
	fmt.Println(a2, s2)

	s3 := a2[:2] //novo slice,mas apontando para o mesmo array a2
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)
}
