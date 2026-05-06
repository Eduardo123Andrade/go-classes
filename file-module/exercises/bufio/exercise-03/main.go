// 3. Lendo um Arquivo Linha a Linha
// Crie um programa que abra um arquivo .txt e imprima apenas as linhas que começam com o caractere # (comentários).
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("../../../files/go-test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			firstChar := line[:1]

			if firstChar == "#" {
				fmt.Println(line)
			}
		}

	}

}
