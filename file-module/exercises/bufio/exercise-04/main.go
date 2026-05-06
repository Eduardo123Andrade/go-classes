package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	const counter = 100_000

	file, err := os.Create("test.txt")

	if err != nil {
		panic("error creating file")
	}

	defer file.Close()
	writer := bufio.NewWriter(file)

	for value := range counter {
		valueToWrite := strconv.Itoa(value + 1)
		writer.WriteString(valueToWrite + "\n")
	}

	writer.Flush()

}
