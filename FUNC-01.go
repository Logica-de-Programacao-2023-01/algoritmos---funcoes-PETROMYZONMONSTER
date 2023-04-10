//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	resultado := media(slice)
	fmt.Println(resultado)
}
func media(slice []int) int {
	resultado := 0
	for _, numeros := range slice {
		resultado += numeros
	}
	return resultado / 5
}
