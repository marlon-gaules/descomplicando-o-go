package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tamanho do arquivo:", stat.Size())
}
