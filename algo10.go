package main

import "fmt"

func main() { //Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	min, max := slice[0], slice[10]
	for _, i := range slice {
		if i < min {
			i = min
		} else if i > max {
			i = max
		}
	}
	fmt.Println("O valor máximo é:", max)
	fmt.Println("O valor mínimo é:", min)
}
