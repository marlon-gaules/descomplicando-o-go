package main

import "fmt"

func main() {
	var x int64 //arquitetura
	x = 10      //8 bytes
	fmt.Println(x)

	//ponteiro de x
	pointer0fX := &x
	fmt.Println(pointer0fX)

	var b bool
	b = false //1 byte
	fmt.Println(b)

	//ponteiro de b
	pointer0fB := &b
	fmt.Println(pointer0fB)

	alfabeto := []string{"a", "b", "c", "d"}
	fmt.Println("alfabeto:", alfabeto)

	primeirasLetras := alfabeto[0:2]
	fmt.Println(primeirasLetras)

	primeirasLetras[0] = "z"
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)

	alfabeto = append(alfabeto, "e")
	primeirasLetras[1] = "y"
	fmt.Println(primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)
}
