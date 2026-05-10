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
