package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxioso interface {
	fazerBaliza()
}

type esportivoLuxioso interface {
	esportivo
	luxioso
	// pode adicionar outros métodos
}

type bmw7 struct {
}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo ...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo baliza ...")
}

func main() {
	var b esportivoLuxioso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
