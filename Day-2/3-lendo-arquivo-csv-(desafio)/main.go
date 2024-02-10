package main

import (
	"bufio"
	"fmt"
	"os"
)

//o arquivo .csv nesse diretório foi baixado desse site https://dados.gov.br/dados/conjuntos-dados/preco-de-medicamentos-no-brasil-consumidor

func main() {

	file, err := os.Open("TA_PRECO_MEDICAMENTO.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
