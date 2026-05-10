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
