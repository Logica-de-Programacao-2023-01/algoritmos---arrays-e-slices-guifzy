package main

import (
	"fmt"
)

func main() { //Faça um programa que leia dois arrays de inteiros de tamanho n e gere um
	// terceiro array que seja a soma dos dois primeiros.
	var n int
	fmt.Println("Digite o tamanho do slice:")
	fmt.Scan(&n)

	lista1 := make([]int, 0, n)
	lista2 := make([]int, 0, n)

	fmt.Println("Digite os valores da primeira lista:")
	for i := 0; len(lista1) < n; i++ {
		fmt.Scan(&i)
		lista1 = append(lista1, i)
	}

	fmt.Println("Digite os valores da segunda lista:")
	for i := 0; len(lista2) < n; i++ {
		fmt.Scan(&i)
		lista2 = append(lista2, i)
	}
	fmt.Println("lista1:", lista1, "lista2:", lista2)

	var soma int
	novaLista := make([]int, 0, n)
	for i := 0; len(novaLista) < n; i++ {
		soma = lista1[i] + lista2[i]
		novaLista = append(novaLista, soma)
	}
	fmt.Println(novaLista)
}
