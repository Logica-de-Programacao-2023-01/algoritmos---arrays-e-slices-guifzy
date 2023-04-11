package main

import "fmt"

func main() { //Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.
	var n int
	fmt.Println("Digite o tamanhho da lista:")
	fmt.Scan(&n)

	lista := make([]int, 0, n)
	fmt.Println("Digite os números da lista:")
	for i := 0; len(lista) < n; i++ {
		fmt.Scan(&i)
		lista = append(lista, i)
	}
	fmt.Println(lista)

	verif := true
	for i := 0; i < n-1; i++ {
		if lista[i] > lista[i+1] {
			verif = false
			break
		}
	}
	if verif == true {
		fmt.Println("Está em ordem crescente.")
	} else {
		fmt.Println("Não está em ordem crescente")
	}
}
