package main

import "fmt"

func main() { //Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
	lista := [4]float64{1.25, 2.3, 3, 4}
	var produto float64
	produto = 1

	for _, num := range lista {
		produto *= num
	}
	fmt.Println(produto)
}
