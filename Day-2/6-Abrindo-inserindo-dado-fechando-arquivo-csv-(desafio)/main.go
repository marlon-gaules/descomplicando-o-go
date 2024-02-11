package main

import (
	"fmt"
	"os"
)

func main() {
	//Adapte o código da aula anterior para escrever uma nova linha qualquer no CSV!

	nova_linha := []byte("Essa linha está reservada para um remédio")
	arquivo, erro := os.OpenFile("TA_PRECO_MEDICAMENTO.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer arquivo.Close()
	arquivo.Write(nova_linha)
	arquivo.WriteString("\n----------\n")

}
