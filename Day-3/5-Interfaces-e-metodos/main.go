package main

import (
	"fmt"
)

type Carro struct {
	Marca string //atributos
	Ano   int
}

type Moto struct {
	Modelo string //atributos
	Ano    int
}

type Veiculo interface {
	Buzinar()
}

func (c *Carro) Buzinar() {
	fmt.Println("fon fon")
}

func (c *Moto) Buzinar() {
	fmt.Println("beep beep")
}

func buzina(v Veiculo) {
	v.Buzinar()
}

func main() {
	carro := Carro{
		Marca: "Audi",
		Ano:   2024,
	}
	fmt.Println(carro)
	carro.Buzinar()

	moto := Moto{
		Modelo: "Honda",
		Ano:    2024,
	}
	fmt.Println(moto)
	moto.Buzinar()

	veiculos := make([]Veiculo, 2)
	veiculos[0] = &carro
	veiculos[1] = &moto
	for i := 0; i < len(veiculos); i++ {
		buzina(veiculos[i])
	}
}
