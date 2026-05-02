// Crie uma função exists(path string) bool que verifique se um arquivo ou diretório realmente existe no disco.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(exists("../../files/file.txt"))
	fmt.Println(exists("file.txt"))
}

func exists(path string) bool {
	file, _ := os.Open(path)

	if file != nil {
		return true
	}

	return false
}
