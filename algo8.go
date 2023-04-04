package main

import "fmt"

func main() { //Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
	slice := []int{1, 1, 3, 4, 5, 6, 6, 8}
	var num int

	fmt.Printf("Informe um número:\n")
	fmt.Scan(&num)

	var novoSlice []int
	x := false
	for _, elemento := range slice {
		if num != elemento {
			x = true
		}
	}

	if x == true {
		slice = append(novoSlice, num)
	}

	fmt.Println(novoSlice)

}
