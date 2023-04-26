package main

import "fmt"

func main() { //Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas
	// os elementos do Array que são maiores que 5. Imprima o novo Slice.
	lista := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	novoSlice := make([]float64, 0, len(lista))

	for _, tamanho := range lista {
		if tamanho > 5 {
			novoSlice = append(novoSlice, tamanho)
		}
	}
	fmt.Println(novoSlice)

}
