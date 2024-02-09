package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	txt := []byte("Esse é o conteúdo do meu arquivo novo")
	var perm fs.FileMode
	perm = 666
	os.WriteFile("arquivo-novo.txt", txt, perm)

	txt2 := []byte("Esse é outro conteúdo do meu arquivo mais novo")
	file, err := os.OpenFile("arquivo-novo2.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	file.Write(txt2)
	file.WriteString("minha string")
	file.WriteAt(txt2, 1)
	file.WriteString("\n----------\n")

	fmt.Fprintf(file, "As permissões foram %d", 666)
}
