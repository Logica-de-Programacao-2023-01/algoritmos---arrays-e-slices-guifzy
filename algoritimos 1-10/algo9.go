package main

import "fmt"

func main() { //Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as posições do Array. Imprima o Array resultante.
	array := [6]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	var num float64
	fmt.Printf("Digite um númeo a ser adicionado em todas as posições da lista:\n")
	fmt.Scan(&num)

	for i := range array {
		array[i] += num
	}

	fmt.Println(array)

}
