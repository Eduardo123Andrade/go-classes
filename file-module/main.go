package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("files/file.txt", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	stats, err := file.Stat()

	if err != nil {
		panic(err)
	}

	mode := stats.Mode()

	fmt.Println(mode)
	fmt.Println(mode.Type())
	fmt.Printf("Binary(mode): %b\n", mode)
}
