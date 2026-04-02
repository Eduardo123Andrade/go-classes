package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World")

	fmt.Println("O programa terminou.")

	// Faz o programa esperar até que você aperte ENTER
	fmt.Println("Pressione ENTER para sair...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
