package main

import "fmt"

func main() { //Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
	lista := [3]int{1, 2, 3}
	var soma int

	for _, num := range lista {
		soma += num
	}
	fmt.Println(soma)
}
