// 1. O Caminho Absoluto
// Escreva uma função que receba um caminho relativo (ex: static/docs) e retorne o caminho absoluto completo no sistema de arquivos.

package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	absPath, err := filepath.Abs("../files/file.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Absolute path:", absPath)
}
