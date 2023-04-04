package main

import "fmt"

func main() { //Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
	lista := []int{1, 2, 3, 4, 5}
	var num int

	fmt.Printf("Digite um número:\n")
	fmt.Scan(&num)
	x := false

	for _, elemento := range lista {
		if num == elemento {
			x = true
			break
		}

	}
	if x == false {
		lista = append(lista, num)
	}
	fmt.Println(lista)

}
