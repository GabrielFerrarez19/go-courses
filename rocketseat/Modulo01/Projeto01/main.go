package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Jogo da Adivinhação")
	println("Um numero aleatorio será sorteado. Tente acertar. O numero é um inteiro entre 0 e 100")

	x := rand.Int64N(101)
	x = 10

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		println("Qual é o seu chute")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			println("O seu chute deve ser um numero inteiro")
			return
		}

		switch {
		case chuteInt < x:
			println("Você errou. O numero sorteado é maior que", chute)
		case chuteInt > x:
			println("Você errou. O numero sorteado é menor que", chute)
		case chuteInt == x:
			fmt.Printf("Parabens o numero era: %d\n"+
				"Você acertou em %d tentativas.\n"+
				"Essas foram as suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt

	}
	fmt.Printf(
		"Infelizmente, Você não acertou o numero, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			" Essas foram as suas tentativas: %v\n",
		x, chutes)
}
