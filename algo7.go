package main

import "fmt"

func main() { //Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
	lista := make([]int, 5)
	var num int

	x := false
	for _, elemento := range lista {

		fmt.Println("digite um número:")
		fmt.Scan(&num)

		if num == elemento {
			x = true
		}

		if x == false {
			lista = append(lista, num)
		} else if x == true {
			break
		}
	}

	fmt.Println(lista)
}
