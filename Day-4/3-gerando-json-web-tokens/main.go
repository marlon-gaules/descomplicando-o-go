package main

// Mão na massa: gerando JSON Web Tokens em Go
// https://jwt.io
// https://github.com/cristalhq/jwt

import (
	"fmt"

	"github.com/marlon-gaules/tokens"
)

func main() {
	fmt.Println("Gerando JWT...")
	tokens.Generate()
	fmt.Println("JWT Gerado com sucesso!")
}
