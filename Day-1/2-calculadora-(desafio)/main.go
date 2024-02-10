package main

import "fmt"

func adicao(a, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func main() {

	// TO DO: Pedir os numeros para o usuário.
	// #desafiodocursodescomplicandoogo.

	var a int
	fmt.Println("Por favor informe um numero int:")
	fmt.Scan(&a)
	fmt.Println("O numero int informado é:", a)

	var b int
	fmt.Println("Por favor informe mais um numero int:")
	fmt.Scan(&b)
	fmt.Println("O numero int informado é:", b)

	resultadoAdicao := adicao(a, b)
	fmt.Println("O resultado da adição dos dois numeros int é:", resultadoAdicao)

	resultadoSubtracao := subtracao(a, b)
	fmt.Println("O resultado da subtração dos dois numeros int é:", resultadoSubtracao)

	resultadoMultiplicacao := multiplicacao(a, b)
	fmt.Println("O resultado da multiplicação dos dois numeros int é:", resultadoMultiplicacao)

	resultadoDivisao := divisao(a, b)
	fmt.Println("O resultado da divisão dos dois numeros int é:", resultadoDivisao)
}
