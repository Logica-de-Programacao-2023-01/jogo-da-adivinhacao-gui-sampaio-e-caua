package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numero int
	var jogarnovamente string
	chances := 1
	lista := []int{}

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for {
		fmt.Print("Escreva um número entre 1 e 100: ")
		fmt.Scan(&numero)

		numeroaleatorio := rand.Intn(100) + 1

		for {
			if numero < numeroaleatorio {
				fmt.Printf("O número é maior que %v\n", numero)
			} else if numero > numeroaleatorio {
				fmt.Printf("O número é menor que %v\n", numero)
			} else if numero == numeroaleatorio {
				fmt.Println("Parabéns, você acertou!")
				fmt.Printf("Você utilizou %v chances\n", chances)
				lista = append(lista, chances)
				chances = 1
				break
			}

			fmt.Print("escreva um número entre 1 e 100: ")
			fmt.Scan(&numero)
			chances++

		}
		fmt.Print("Você quer jogar novamente? (sim/nao):")
		fmt.Scan(&jogarnovamente)

		if jogarnovamente != "sim" && jogarnovamente != "nao" {
			fmt.Print("Você quer jogar novamente? (sim/nao):")
			fmt.Scan(&jogarnovamente)
		}

		if jogarnovamente == "nao" {
			for i := 0; i < len(lista); i++ {
				fmt.Printf("Na rodada %v você utilizou %v chances\n", i+1, lista[i])
			}
			break
		}

	}
}
