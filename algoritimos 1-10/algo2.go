package main

import "fmt"

func main() { //Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.
	lista := []int{1, 2, 3, 4, 5}
	lista = append(lista[:2], lista[3:]...)
	fmt.Println(lista)
}
