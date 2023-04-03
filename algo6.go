package main

import "fmt"

func main() { //Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca.
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var num int

	x := false
	fmt.Println("Digite um número")
	fmt.Scan(&num)
	for _, elemento := range lista {
		if num == elemento {
			x = true
		}
	}

	if x == true {
		fmt.Println("Seu número está na lista.")
	} else {
		fmt.Println("Seu número não está na lista")
	}
}
