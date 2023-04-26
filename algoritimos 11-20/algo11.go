package main

import "fmt"

func main() { //Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.
	matriz := [2][3]int{}

	for l := 0; l < 2; l++ {
		for c := 0; c < 3; c++ {
			fmt.Printf("Digite um número [%d][%d]:", l, c)
			fmt.Scan(&matriz[l][c])
		}
	}

	fmt.Println("Seu array bidimensional:")
	for l := 0; l < 2; l++ {
		for c := 0; c < 3; c++ {
			fmt.Printf("%d", matriz[l][c])
		}
		fmt.Println()
	}

	var linha, coluna int

	fmt.Printf("Informe o índice de linha: %d\n", linha)
	fmt.Scan(&linha)
	fmt.Printf("Informe o índice de coluna: %d\n", coluna)
	fmt.Scan(&coluna)

	fmt.Println("O seu numero é:", matriz[linha][coluna])
}
