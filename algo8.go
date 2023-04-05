package main

import "fmt"

func main() { //Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
	slice := []string{"gui", "lucas", "felipe", "maria", "leo", "luisa", "bruno", "jao"}
	var num string

	fmt.Printf("Informe o número a ser removido da lista %v:\n", slice)
	fmt.Scan(&num)

	novoSlice := make([]string, 0, len(slice))
	for _, elemento := range slice {
		if elemento != num {
			novoSlice = append(novoSlice, elemento)
		}
	}
	fmt.Println(novoSlice)
}
