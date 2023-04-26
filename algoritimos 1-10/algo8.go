package main

import "fmt"

func main() { //Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
	slice := []string{"gui", "lucas", "felipe", "maria", "leo", "luisa", "bruno", "jao"}
	var nome string

	fmt.Printf("Informe o nome a ser removido da lista: %v\n", slice)
	fmt.Scan(&nome)

	novoSlice := make([]string, 0, len(slice))
	for _, elemento := range slice {
		if elemento != nome {
			novoSlice = append(novoSlice, elemento)
		}
	}
	fmt.Println(novoSlice)
}
