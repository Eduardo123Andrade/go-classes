// Liste todos os arquivos (apenas na pasta atual, sem recursão) que possuam a extensão .txt.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.ReadDir("../../files")

	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		ext := filepath.Ext(file.Name())
		if ext == ".txt" {
			fmt.Println(file.Name())
		}
	}
}
