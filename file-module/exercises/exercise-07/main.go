// Implemente uma função que percorra uma pasta e todas as suas subpastas, imprimindo o caminho de cada arquivo encontrado que seja maior que 1MB.

package main

import (
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("../../files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Size() > 1024*1024 {
			println(path)
		}

		return nil
	})

}
