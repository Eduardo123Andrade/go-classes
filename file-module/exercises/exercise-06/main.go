// Escreva um programa que receba o nome de um arquivo e verifique se o usuário atual tem permissão de escrita.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("../../files/file.txt", os.O_WRONLY, 0)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("O usuário tem permissão de escrita no arquivo.")
}
