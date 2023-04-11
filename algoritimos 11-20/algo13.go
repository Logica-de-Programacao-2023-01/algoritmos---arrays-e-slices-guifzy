package main

import "fmt"

func main() { //Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	var num int
	fmt.Println("Digite um número a ser adicionado ao primeiro e último elemento da lista:")
	fmt.Scan(&num)

	array[0] += num
	array[6] += num

	fmt.Println(array)

}