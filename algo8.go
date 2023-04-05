package main

import "fmt"

func main() { //Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
	slice := []int{1, 1, 3, 4, 5, 6, 6, 8}
	var num int

	fmt.Printf("Informe o número a ser removido da lista %d:\n", slice)
	fmt.Scan(&num)

	novoSlice := make([]int, 0, len(slice))
	for _, elemento := range slice {
		if num != elemento {
			novoSlice = append(novoSlice, elemento)
		}
	}

	fmt.Println(novoSlice)

}
