package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.ReadFile("arquivo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Alterando array de dados para string:", string(file))

	file2, err := os.Open("arquivo2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	file3, err := os.Open("arquivo3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file3)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		fmt.Println(b)
		fmt.Println(string(b))
	}
}
