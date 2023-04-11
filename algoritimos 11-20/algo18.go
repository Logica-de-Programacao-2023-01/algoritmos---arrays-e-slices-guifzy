package main

import (
	"fmt"
)

func main() { //Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.
	var n int
	fmt.Println("Digite um número inteiro positivo:")
	fmt.Scan(&n)

	lista := make([]int, 0, n)
	for i := 2; len(lista) < n; i++ {
		primo := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo == true {
			lista = append(lista, i)
		}
	}
	fmt.Println(lista)

}
