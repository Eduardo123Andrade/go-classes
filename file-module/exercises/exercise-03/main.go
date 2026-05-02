// Escreva um código que crie a estrutura de pastas projeto/src/internal/config de uma só vez, garantindo que não dê erro caso as pastas já existam.

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	err := os.MkdirAll("project/src/internal/config", fs.ModePerm)

	if err != nil {
		panic(err)
	}

	fmt.Println("Directory created successfully or already exists.")
}
