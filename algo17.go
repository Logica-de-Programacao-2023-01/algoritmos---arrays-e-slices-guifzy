package main

import "fmt"

func main() { //Crie um Array de inteiros com 10 elementos.
	// Calcule e imprima a soma dos elementos nas posições pares do Array.
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var soma int

	for _, i := range lista {
		if i%2 == 0 {
			soma += i
		}
	}
	fmt.Println(soma)
}
