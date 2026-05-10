// 8. Lendo até um Caractere Específico
// Use bufio.Reader e o método ReadString ou ReadBytes para ler um arquivo até encontrar o caractere @.
// Capture o que foi lido e continue a leitura até o próximo @.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("../../../files/file2.txt", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		content, err := reader.ReadString('@')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		cleanContent := content[:len(content)-1]

		fmt.Printf("%s\n", cleanContent)

	}

}
