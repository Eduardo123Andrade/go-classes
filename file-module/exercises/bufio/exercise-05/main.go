// 5. Prefixador de Linhas
// Leia um arquivo de entrada e gere um arquivo de saída onde cada linha recebe um prefixo com o seu número
// (Ex: 1: Primeira linha, 2: Segunda linha). Use Scanner para ler e bufio.Writer para escrever.
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("../../../files/file.txt")

	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	outPutFile, err := os.Create("test-copy.txt")

	if err != nil {
		panic(err)
	}

	defer outPutFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outPutFile)

	numberPrefix := 0
	for scanner.Scan() {
		text := scanner.Text()
		numberPrefix++
		valueToWrite := strconv.Itoa(numberPrefix) + " - " + text + "\n"
		writer.WriteString(valueToWrite)
	}

	writer.Flush()

}
