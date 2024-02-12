package main

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

func percorreDiretorio(path string, d fs.DirEntry, err error) error {
	fmt.Println("É diretório? ", d.IsDir())
	info, _ := d.Info()
	fmt.Println("Name:: ", info.Name())
	fmt.Println("----------\n")
	return nil
}

func main() {
	filepath.WalkDir(".", percorreDiretorio)

	percorreOutrosDiretorio := path.Join("teste", "teste2")
	fmt.Println(percorreOutrosDiretorio)
}
