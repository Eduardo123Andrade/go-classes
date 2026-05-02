// Dado o caminho files/file.txt, seu programa deve imprimir separadamente:
// O diretório pai.
// O nome do arquivo com extensão.
// A extensão do arquivo.
// O nome do arquivo sem a extensão.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("../../files/file.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	path := file.Name()

	ext := filepath.Ext(path)
	fileName := filepath.Base(path)
	name := fileName[:len(fileName)-len(ext)]

	parentPath := filepath.Dir(path)
	parent := filepath.Base(parentPath)

	println("Diretório pai:", parent)
	println("Nome do arquivo com extensão:", fileName)
	println("Extensão do arquivo:", ext)
	fmt.Println("Nome do arquivo sem a extensão:", name)
}
