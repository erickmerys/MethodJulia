package main

import (
	"errors"
	"fmt"
)

func calcularRaizQuadrada(n int) (int, error) {
	var resultado int
	var raiz int
	var i int
	for i = 1; i < n; i++ {

		sucessor := i + 1
		raiz = i * i
		resultado = raiz + i + sucessor

		if resultado == n {
			return sucessor, nil
		}else if resultado > n{
			return 0, errors.New("não é possível calcular raiz quadrada não exata")
		}
	}
	return 0, errors.New("não é possível calcular raiz quadrada não exata")
	
}

func main() {
	var num int

	/* fmt.Print("Digite um numero para calcular a raiz quadrada: ")
	fmt.Scan(&num) */
	num = 1023

	if raiz, err := calcularRaizQuadrada(num); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("A raiz quadrada de %d é: %d\n", num , raiz)
	}
}
