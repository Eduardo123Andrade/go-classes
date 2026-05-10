// 10. Inversor de Arquivos (Buffer de trás para frente)
// Implemente um programa que leia um arquivo e escreva um novo arquivo com as mesmas linhas, mas na ordem inversa (a última linha do original vira a primeira do novo).

// Dica: Você provavelmente precisará ler o arquivo para uma slice em memória usando Scanner e depois descarregar usando Writer.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("../../../files/file.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	content := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		content = append(content, text)
	}

	newfile, err := os.OpenFile("reverse-file.txt", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer newfile.Close()

	writer := bufio.NewWriter(newfile)

	for i := len(content) - 1; i >= 0; i-- {
		_, err := writer.WriteString(content[i] + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	writer.Flush()

}
