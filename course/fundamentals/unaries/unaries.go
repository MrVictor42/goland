package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix em go
	x++ // x = x + 1 ou x += 1
	fmt.Println(x)
	
	y -- // y = y - 1 ou y -= 1
	fmt.Println(y)
}