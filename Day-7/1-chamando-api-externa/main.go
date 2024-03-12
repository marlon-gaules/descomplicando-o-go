package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonGetResponse struct {
	Count    int64
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}

func main() {
	url := "https://pokeapi.co/api/v2/pokemon"
	fmt.Println("Acessando a API na url ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("O status code retornado pela url é", resp.StatusCode)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	pokemonGetResponse := new(PokemonGetResponse)
	err = json.Unmarshal(body, pokemonGetResponse)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("O nome dos 9 primeiros pokemons retornados são: ")
	for i := 0; i < 9; i++ {
		fmt.Println(pokemonGetResponse.Results[i].Name)
	}
}
