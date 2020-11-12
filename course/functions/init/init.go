package main

import "fmt"

// não se pode ter duas funções main dentro do mesmo pacote, mas pode ter mais de um init por pacote

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main....")
}
