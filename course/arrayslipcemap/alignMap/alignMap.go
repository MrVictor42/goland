package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12532.42,
			"Guga Pereira":   14242,
		},
		"J": {
			"Jéssica Lopes": 23000,
		},
		"P": {
			"Pedro Henrique": 14222,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
