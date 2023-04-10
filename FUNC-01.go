//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

package main

import (
	"fmt"
)

func main() {
	var slice = []int{}
	var n, x, soma int
	fmt.Println("Digite o tamanho da slice:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Digite os elementos da slice:")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma += slice[i]
	}
	resultado := media(soma, n)
	fmt.Println(resultado)
}

func media(soma int, n int) int {
	resultado := soma / n
	return resultado
}
