// Crie um script que percorra um diretório recursivamente e conte quantas linhas totais de código existem em todos os arquivos .go.

package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	directoryPath := "../../"

	var lineCount int = 0

	filepath.WalkDir(directoryPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".go" {
			if filepath.Ext(path) != ".go" {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				lineCount++
			}
		}

		return nil
	})

	fmt.Println("Total lines in Go files:", lineCount)
}
