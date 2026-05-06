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
