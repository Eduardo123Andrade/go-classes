// 7. O Método Peek
// Use um bufio.Reader para espiar os primeiros 5 bytes de um arquivo sem "consumi-los"
// (ou seja, sem mover o ponteiro de leitura). Depois, leia o arquivo inteiro normalmente e
// verifique se esses 5 bytes ainda aparecem no início da leitura.

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

	reader := bufio.NewReader(file)
	content, err := reader.Peek(5)

	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	text := scanner.Text()

	firstValues := text[:len(content)]

	if strValue := string(content); strValue == firstValues {
		fmt.Println("same", strValue)
	}
}
