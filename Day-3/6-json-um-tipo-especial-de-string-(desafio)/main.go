package main

import (
	"encoding/json"
	"fmt"
)

type Carro struct {
	Marca string `json:"marca"` // use tags para mudar o nome do campo
	Ano   int
}

type Moto struct {
	Marca string
	Ano   int
}

type Veiculo interface {
}

func main() {
	carro := Carro{
		Marca: "Audi",
		Ano:   2024,
	}

	moto := Moto{
		Marca: "Honda",
		Ano:   2024,
	}

	veiculos := []Veiculo{&carro, &moto}

	changingVeiculosToJson, err := json.Marshal(veiculos)
	if err != nil {
		fmt.Println("Error ::", err.Error())
	}
	fmt.Printf("%s\n", changingVeiculosToJson)

	fmt.Println(carro, moto) //retorno sem alterar para Json

	//alterando de json para estrura default Go
	s := "[{\"Marca\":\"Audi\",\"Ano\":2024}]" //formato Json
	carros := []Carro{}
	err = json.Unmarshal([]byte(s), &carros)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carros)
}
