// 1. Eco do Teclado
// Crie um programa que leia linhas digitadas pelo usuário no terminal (os.Stdin) e as imprima em letras maiúsculas até que o usuário digite "sair".

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	exit := false

	scanner := bufio.NewScanner(os.Stdin)
	for !exit {
		fmt.Printf("Type something: ")

		if !scanner.Scan() {
			panic("scanner not open")

		}
		value := scanner.Text()
		// fmt.Printf("You typed: %s", value)

		if value == "sair" {
			exit = true
			return 
		}

		upperValue := strings.ToUpper(value)
		fmt.Println("You typed: ", upperValue)
	}
}