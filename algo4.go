package main

import "fmt"

func main() { //Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
	var larg int
	fmt.Println("Digite a largura da lista:")
	fmt.Scan(&larg)

	lista := make([]int, larg)

	for i := 0; i < larg; i++ {
		fmt.Println("Digite os números da lista:")
		fmt.Scan(&lista[i])
	}

	var soma int

	for _, num := range lista {
		soma += num

	}
	fmt.Println("A soma dos valores da lista é:", soma)
}
