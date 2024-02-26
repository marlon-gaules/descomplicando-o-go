package main

import (
	"fmt"
)

//struct não anônima
type Carro2 struct {
	Marca string //atributos
	Ano   int
}

func main() {

	//Array
	caracteres := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(caracteres)
	fmt.Println(len(caracteres))
	for i := 0; i < len(caracteres); i++ {
		fmt.Println(caracteres[i])
	}

	matriz_caracteres := [2][5]string{
		{"a", "b", "c", "d", "e"},
		{"f", "g", "h", "i", "j"},
	}
	for linha := 0; linha < len(matriz_caracteres); linha++ {
		for coluna := 0; coluna < len(matriz_caracteres[linha]); coluna++ {
			fmt.Println(matriz_caracteres[linha][coluna])
		}
	}

	matriz_caracteres2 := [3][3]string{
		{"a", "b", "c"},
		{"1", "2", "3"},
		{"I", "II", "III"},
	}
	for linha := 0; linha < len(matriz_caracteres2); linha++ {
		for coluna := 0; coluna < len(matriz_caracteres2[linha]); coluna++ {
			fmt.Println(matriz_caracteres2[linha][coluna])
		}
	}

	//Slice
	caracteres2 := []string{"a", "b", "c", "d", "e"}
	caracteres2 = append(caracteres2, "f", "g")
	fmt.Println(caracteres2)

	capacidade := 200
	tamanhoDoSlice := 10
	caracteres3 := make([]string, tamanhoDoSlice, capacidade)
	caracteres3[4] = "4"
	caracteres3 = append(caracteres3, "5")
	fmt.Println(caracteres3)

	//Map
	alfabeto := make(map[string]string)
	alfabeto["vogais"] = "aeiou"
	alfabeto["consoantes"] = "bcdfg..."
	fmt.Println(alfabeto)

	alfabeto2 := map[string]string{
		"vogais":     "aeiou",
		"consoantes": "bcdfg...",
	}
	fmt.Println(alfabeto2)

	alfabeto3 := map[int]string{
		0: "aeiou",
		1: "bcdfg...",
	}
	fmt.Println(alfabeto3)
	fmt.Println(alfabeto3[0])

	//Estruturas
	//struct anônima
	carro := struct {
		Marca string
		Ano   int
	}{
		Marca: "DMC DeLorean",
		Ano:   1982,
	}
	fmt.Println(carro)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Ano)

	carro2 := Carro2{
		Marca: "Fiat",
		Ano:   2024,
	}
	fmt.Println(carro2)
	fmt.Println(carro2.Marca)
	fmt.Println(carro2.Ano)
}
