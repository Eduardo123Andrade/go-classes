package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main(){
	defer func(){
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("O jogo acabou!")
		fmt.Println("Pressione ENTER para sair...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}()
	
	const MAX_ATTENPTS = 10

	fmt.Println("Jogo da adivinhacao")
	fmt.Println("Um número será sorteado. Tente acertar. O número é um inteiro entre 1 e 100.")
	fmt.Printf("Você terá ate %d tentativas\n", MAX_ATTENPTS)
	fmt.Println("--------------------------------------------------------------------------------")

	randomNumber := rand.Int64N(101)
	kicks := [MAX_ATTENPTS]int64{}
	lose := true

	for i := range MAX_ATTENPTS {
		fmt.Print("Digite o número: ")
		var kick int64
		fmt.Scan(&kick)

		if kick < randomNumber {
			fmt.Printf("O número sorteado é maior. %d tentativas restantes\n", MAX_ATTENPTS - (i + 1))
		} else if kick > randomNumber {
			fmt.Printf("O número sorteado é menor. %d tentativas restantes\n", MAX_ATTENPTS - (i + 1))
		} else {
			fmt.Printf(
				"Parabéns! Você acertou! O número era: %d\n"+
					"Você acertou em %d tentativas\n"+
					"Essas foram as suas tentativas: %v\n",
				randomNumber, i+1, kicks[:i],
			)
			lose = false
			break
		}

		kicks[i] = kick
	}

	if lose {
		fmt.Printf(
			"Infelizmente, você não acertou o número, que era: %d\n"+
				"Você teve %d tentativas.\n"+
				"Essas foram as suas tentativas: %v\n",
			randomNumber, MAX_ATTENPTS, kicks,
		)
	}
}