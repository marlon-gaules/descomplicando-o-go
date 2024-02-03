package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	b, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(string(b))
}
