package main

import "fmt"

func main() { //Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
	array := [5]int{3, 6, 5, 12, 10}

	slice := []int{}
	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}
	fmt.Println(slice)

}
