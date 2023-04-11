package main

import "fmt"

func main() { //Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de
	// elementos que deverão ser trocados de posição. Imprima o Slice resultante.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Slice atual:", slice)
	fmt.Println("Digite os índices que trocarão de lugar:")

	var a, b int
	fmt.Println("Digite o primeiro índice:")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo índice:")
	fmt.Scan(&b)

	if a != b {
		slice[a], slice[b] = slice[b], slice[a]
		fmt.Println("Novo sliice:", slice)
	} else {
		fmt.Println("Digite índices diferentes.")
	}

}
