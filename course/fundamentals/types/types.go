package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// O GO NÃO TEM CHAR. O CHAR AQUI É DO TIPO INT32

	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000))

	// sem sinal (só positivos)
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	// com sinal
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é:", i1)
	fmt.Println("o tipo da variavel i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode
	fmt.Println("O rune é:", i2)

	//números reais
	var x float32 = 49.99
	fmt.Println("o tipo de x é:", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é:", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("o tipo do bo é:", reflect.TypeOf(bo))
	fmt.Println(bo)
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Victor"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string s1 é:", len(s1))

	//outro tipo de string
	s2 := `Olá meu nome é Victor`
	fmt.Println("o tamanho da string s2 é:", len(s2))
}
