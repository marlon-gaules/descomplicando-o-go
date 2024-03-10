package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var CurrentVersion = "desenvolvimento"

//colocando flag de tempo de execução com pacote flag do Go
var packages = flag.String("packages", "kubectl", "Pacotes a serem instalados, separados por vírgula. valores possíveis: kubectl")

func main() {
	//garante o que estou fazendo está pegando as flags
	flag.Parse()

	fmt.Println("Tentando instalar", *packages)
	fmt.Println("Versão", CurrentVersion)

	//fazendo um split da minha string baseado nas minhas vírgulas
	packagesArray := strings.Split(*packages, ",")

	//percorrendo meu array
	for _, p := range packagesArray {
		switch p {
		case "kubectl":
			{
				fmt.Println("Instalando kubectl...")
				command, args := GetKubeCTLInstallCommand()
				cmd := exec.Command(command, args...)

				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}
			}
		default:
			fmt.Println("Pacote", p, "não suportado pelo nosso script!")
		}
	}
}
